package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/app/gateway"
	"github.com/tsundata/assistant/internal/pkg/utils"
	slackVendor "github.com/tsundata/assistant/internal/pkg/vendors/slack"
	"github.com/valyala/fasthttp"
	"go.uber.org/zap"
	"html/template"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type GatewayController struct {
	opt       *gateway.Options
	rdb       *redis.Client
	logger    *zap.Logger
	subClient pb.SubscribeClient
	msgClient pb.MessageClient
}

func NewGatewayController(opt *gateway.Options, rdb *redis.Client, logger *zap.Logger, subClient pb.SubscribeClient, msgClient pb.MessageClient) *GatewayController {
	return &GatewayController{opt: opt, rdb: rdb, logger: logger, subClient: subClient, msgClient: msgClient}
}

func (gc *GatewayController) Index(c *fasthttp.RequestCtx) {
	c.Response.SetBody(utils.StringToByte("Gateway"))
}

func (gc *GatewayController) Apps(c *fasthttp.RequestCtx) {
	c.Response.Header.Set("Content-Type", "text/html; charset=utf-8")
	t := template.Must(template.New("").Parse(`<table>{{range .}}<tr><td>{{.}}</td></tr>{{end}}</table>`))
	names := []string{"slack", "email"}
	if err := t.Execute(c.Response.BodyWriter(), names); err != nil {
		gc.logger.Error(err.Error())
	}
}

func (gc *GatewayController) SlackShortcut(c *fasthttp.RequestCtx) {
	// verificationTokens
	s, err := slackVendor.SlashShortcutParse(&c.Request)
	if err != nil {
		gc.logger.Error(err.Error())
		return
	}

	if !s.ValidateToken(gc.opt.Verification) {
		gc.logger.Info("unvalidated verificationTokens")
		return
	}

	if s.Type == "shortcut" {
		switch s.CallbackID {
		case "report":
			gc.logger.Info("report")
		}
	}

	if s.Type == "message_action" {
		switch s.CallbackID {
		case "delete":
			gc.logger.Info("delete")
		case "run":
			reply, err := gc.msgClient.Run(context.Background(), &pb.MessageRequest{
				Text: s.Message.Text,
			})
			if err != nil {
				gc.logger.Error(err.Error())
				return
			}
			err = slackVendor.ResponseText(s.ResponseURL, reply.GetText())
			if err != nil {
				gc.logger.Error(err.Error())
				return
			}
		}
	}

	c.Response.SetBodyString("OK")
}

func (gc *GatewayController) SlackCommand(c *fasthttp.RequestCtx) {
	// verificationTokens
	s, err := slackVendor.SlashCommandParse(&c.Request)
	if err != nil {
		gc.logger.Error(err.Error())
		c.Error(err.Error(), http.StatusBadRequest)
		return
	}

	if !s.ValidateToken(gc.opt.Verification) {
		gc.logger.Info("unvalidated verificationTokens")
		c.Error("unvalidated verificationTokens", http.StatusBadRequest)
		return
	}

	// parse
	switch s.Command {
	case "/view":
		id, err := strconv.Atoi(s.Text)
		if err != nil {
			gc.logger.Error(err.Error())
			c.Error(err.Error(), http.StatusBadRequest)
			return
		}

		reply, err := gc.msgClient.Get(context.Background(), &pb.MessageRequest{
			Id: int64(id),
		})
		if err != nil {
			gc.logger.Error(err.Error())
			c.Error(err.Error(), http.StatusBadRequest)
			return
		}

		if reply.GetText() != "" {
			err = slackVendor.ResponseText(s.ResponseURL, reply.GetText())
			if err != nil {
				gc.logger.Error(err.Error())
				c.Error(err.Error(), http.StatusBadRequest)
				return
			}
		} else {
			err = slackVendor.ResponseText(s.ResponseURL, "view failed")
			if err != nil {
				gc.logger.Error(err.Error())
				c.Error(err.Error(), http.StatusBadRequest)
				return
			}
		}
	case "/run":
		id, err := strconv.Atoi(s.Text)
		if err != nil {
			gc.logger.Error(err.Error())
			c.Error(err.Error(), http.StatusBadRequest)
			return
		}
		reply, err := gc.msgClient.Get(context.Background(), &pb.MessageRequest{
			Id: int64(id),
		})
		if err != nil {
			gc.logger.Error(err.Error())
			c.Error(err.Error(), http.StatusBadRequest)
			return
		}

		if reply.GetText() != "" {
			r, err := gc.msgClient.Run(context.Background(), &pb.MessageRequest{
				Text: reply.GetText(),
			})
			if err != nil {
				gc.logger.Error(err.Error())
				c.Error(err.Error(), http.StatusBadRequest)
				return
			}
			err = slackVendor.ResponseText(s.ResponseURL, r.GetText())
			if err != nil {
				gc.logger.Error(err.Error())
				c.Error(err.Error(), http.StatusBadRequest)
				return
			}
		}
	}

	c.Response.SetBodyString("OK")
}

func (gc *GatewayController) SlackEvent(c *fasthttp.RequestCtx) {
	body := c.Request.Body()

	api := slack.New(gc.opt.Token)
	eventsAPIEvent, err := slackevents.ParseEvent(body, slackevents.OptionVerifyToken(&slackevents.TokenComparator{VerificationToken: gc.opt.Verification}))
	if err != nil {
		gc.logger.Error(err.Error())
		return
	}

	if eventsAPIEvent.Type == slackevents.URLVerification {
		var r *slackevents.ChallengeResponse
		err := json.Unmarshal(body, &r)
		if err != nil {
			gc.logger.Error(err.Error())
			c.Error(err.Error(), http.StatusBadRequest)
			return
		}
		c.Response.SetBodyString(r.Challenge)
		return
	}

	if eventsAPIEvent.Type == slackevents.CallbackEvent {
		innerEvent := eventsAPIEvent.InnerEvent
		switch ev := innerEvent.Data.(type) {
		case *slackevents.MessageEvent:
			// ignore bot message
			if ev.ClientMsgID != "" {
				// ignore repeated message
				rKey := "message:repeated:" + ev.ClientMsgID
				isRepeated := gc.rdb.Get(context.Background(), rKey).Val()
				if len(isRepeated) > 0 {
					return
				}
				gc.rdb.Set(context.Background(), rKey, time.Now().Unix(), 7*24*time.Hour)

				// special <url>, utf8 whitespace
				re := regexp.MustCompile("<" + utils.UrlRegex + ">")
				urls := re.FindAllString(ev.Text, -1)
				for _, url := range urls {
					ev.Text = strings.ReplaceAll(ev.Text, url, strings.TrimRight(strings.TrimLeft(url, "<"), ">"))
				}
				re = regexp.MustCompile(`[\s\p{Zs}]+`)
				ev.Text = re.ReplaceAllString(ev.Text, " ")

				reply, err := gc.msgClient.Create(context.Background(), &pb.MessageRequest{
					Uuid: ev.ClientMsgID,
					Text: ev.Text,
				})
				if err != nil {
					gc.logger.Error(err.Error())
					return
				}

				if reply.GetId() > 0 {
					_, _, err = api.PostMessage(ev.Channel, slack.MsgOptionText(fmt.Sprintf("ID: %d", reply.GetId()), false))
				} else {
					for _, item := range reply.GetText() {
						_, _, err = api.PostMessage(ev.Channel, slack.MsgOptionText(item, false))
					}
				}
				if err != nil {
					gc.logger.Error(err.Error())
					return
				}
			}
		}
	}

	c.Response.SetBodyString("OK")
}
