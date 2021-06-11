package listener

import (
	"encoding/json"
	"fmt"
	"github.com/nats-io/nats.go"
	"github.com/tsundata/assistant/internal/pkg/config"
	"github.com/tsundata/assistant/internal/pkg/event"
	"github.com/tsundata/assistant/internal/pkg/model"
	"github.com/tsundata/assistant/internal/pkg/transports/http"
	"github.com/tsundata/assistant/internal/pkg/utils"
	"github.com/valyala/fasthttp"
	"log"
)

func RegisterEventHandler(bus *event.Bus, config *config.AppConfig) error {
	err := bus.Subscribe(event.EchoSubject, func(msg *nats.Msg) {
		fmt.Println(msg)
	})
	if err != nil {
		return err
	}

	err = bus.Subscribe(event.SendMessageSubject, func(msg *nats.Msg) {
		var message model.Message
		err := json.Unmarshal(msg.Data, &message)
		if err != nil {
			log.Println(err)
			return
		}

		client := http.NewClient()
		resp, err := client.PostJSON(config.Slack.Webhook, map[string]interface{}{
			"text": message.Text,
		})
		if err != nil {
			log.Println(err)
			return
		}

		_ = utils.ByteToString(resp.Body())
		fasthttp.ReleaseResponse(resp)
	})
	if err != nil {
		return err
	}

	return nil
}
