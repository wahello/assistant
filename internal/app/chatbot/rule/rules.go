package rule

import (
	"context"
	"errors"
	"fmt"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/rulebot"
	"github.com/tsundata/assistant/internal/pkg/util"
	"github.com/tsundata/assistant/internal/pkg/version"
	"io"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

var rules = []Rule{
	{
		Regex: `version`,
		Help:  `Version info`,
		Parse: func(ctx context.Context, comp rulebot.IComponent, s string, args []string) []string {
			return []string{version.Info()}
		},
	},
	{
		Regex: `menu`,
		Help:  `Show menu`,
		Parse: func(ctx context.Context, comp rulebot.IComponent, s string, args []string) []string {
			if comp.Middle() == nil {
				return []string{"empty client"}
			}
			reply, err := comp.Middle().GetMenu(ctx, &pb.TextRequest{})
			if err != nil {
				return []string{"error call: " + err.Error()}
			}

			if reply.GetText() == "" {
				return []string{"empty menu"}
			}

			return []string{reply.GetText()}
		},
	},
	{
		Regex: `qr\s+(.*)`,
		Help:  `Generate QR code`,
		Parse: func(ctx context.Context, comp rulebot.IComponent, s string, args []string) []string {
			if comp.Middle() == nil {
				return []string{"empty client"}
			}
			if len(args) != 2 {
				return []string{"error args"}
			}

			txt := args[1]
			reply, err := comp.Middle().GetQrUrl(ctx, &pb.TextRequest{
				Text: txt,
			})
			if err != nil {
				return []string{"error call: " + err.Error()}
			}

			return []string{
				reply.GetText(),
			}
		},
	},
	{
		Regex: `ut\s+(\d+)`,
		Help:  `Unix Timestamp`,
		Parse: func(ctx context.Context, comp rulebot.IComponent, s string, args []string) []string {
			if len(args) != 2 {
				return []string{"error args"}
			}

			tt, err := strconv.ParseInt(args[1], 10, 64)
			if err != nil {
				return []string{"error call: " + err.Error()}
			}

			t := time.Unix(tt, 0)

			return []string{
				t.String(),
			}
		},
	},
	{
		Regex: `rand\s+(\d+)\s+(\d+)`,
		Help:  `Unix Timestamp`,
		Parse: func(ctx context.Context, comp rulebot.IComponent, s string, args []string) []string {
			if len(args) != 3 {
				return []string{"error args"}
			}

			minArg := args[1]
			maxArg := args[2]
			min, err := strconv.Atoi(minArg)
			if err != nil {
				return []string{"error call: " + err.Error()}
			}
			max, err := strconv.Atoi(maxArg)
			if err != nil {
				return []string{"error call: " + err.Error()}
			}

			rand.Seed(time.Now().Unix())
			t := rand.Intn(max-min) + min

			return []string{
				strconv.Itoa(t),
			}
		},
	},
	{
		Regex: `pwd\s+(\d+)`,
		Help:  `Generate Password`,
		Parse: func(ctx context.Context, comp rulebot.IComponent, s string, args []string) []string {
			if len(args) != 2 {
				return []string{"error args"}
			}

			lenArg := args[1]
			length, err := strconv.Atoi(lenArg)
			if err != nil {
				return []string{"error call: " + err.Error()}
			}

			pwd := util.GeneratePassword(length, "lowercase|uppercase|numbers")

			return []string{
				pwd,
			}
		},
	},
	{
		Regex: `subs\s+list`,
		Help:  `List subscribe`,
		Parse: func(ctx context.Context, comp rulebot.IComponent, s string, args []string) []string {
			if comp.Subscribe() == nil {
				return []string{"empty client"}
			}
			reply, err := comp.Subscribe().List(ctx, &pb.SubscribeRequest{})
			if err != nil {
				return []string{"error call: " + err.Error()}
			}

			if reply.GetText() == nil {
				return []string{"empty subscript"}
			}

			return reply.GetText()
		},
	},
	{
		Regex: `subs\s+open\s+(.*)`,
		Help:  `Open subscribe`,
		Parse: func(ctx context.Context, comp rulebot.IComponent, s string, args []string) []string {
			if comp.Subscribe() == nil {
				return []string{"empty client"}
			}
			if len(args) != 2 {
				return []string{"error args"}
			}

			reply, err := comp.Subscribe().Open(ctx, &pb.SubscribeRequest{
				Text: args[1],
			})
			if err != nil {
				return []string{"error call: " + err.Error()}
			}
			if reply.GetState() {
				return []string{"ok"}
			}

			return []string{"failed"}
		},
	},
	{
		Regex: `subs\s+close\s+(.*)`,
		Help:  `Close subscribe`,
		Parse: func(ctx context.Context, comp rulebot.IComponent, s string, args []string) []string {
			if comp.Subscribe() == nil {
				return []string{"empty client"}
			}
			if len(args) != 2 {
				return []string{"error args"}
			}

			reply, err := comp.Subscribe().Close(ctx, &pb.SubscribeRequest{
				Text: args[1],
			})
			if err != nil {
				return []string{"error call: " + err.Error()}
			}
			if reply.GetState() {
				return []string{"ok"}
			}

			return []string{"failed"}
		},
	},
	{
		Regex: `view\s+(\d+)`,
		Help:  `View message`,
		Parse: func(ctx context.Context, comp rulebot.IComponent, s string, args []string) []string {
			if comp.Message() == nil {
				return []string{"empty client"}
			}
			if len(args) != 2 {
				return []string{"error args"}
			}

			id, err := strconv.ParseInt(args[1], 10, 64)
			if err != nil {
				return []string{"error args"}
			}
			messageReply, err := comp.Message().Get(ctx, &pb.MessageRequest{Id: id})
			if err != nil {
				return []string{"error call: " + err.Error()}
			}

			if messageReply.Id == 0 {
				return []string{"no message"}
			}

			return []string{messageReply.GetText()}
		},
	},
	{
		Regex: `run\s+(\d+)`,
		Help:  `Run message`,
		Parse: func(ctx context.Context, comp rulebot.IComponent, s string, args []string) []string {
			if comp.Message() == nil {
				return []string{"empty client"}
			}
			if len(args) != 2 {
				return []string{"error args"}
			}

			id, err := strconv.ParseInt(args[1], 10, 64)
			if err != nil {
				return []string{"error args"}
			}

			reply, err := comp.Message().Run(ctx, &pb.MessageRequest{Id: id})
			if err != nil {
				return []string{"error call: " + err.Error()}
			}

			return []string{reply.GetText()}
		},
	},
	{
		Regex: `doc`,
		Help:  `Show action docs`,
		Parse: func(ctx context.Context, comp rulebot.IComponent, s string, args []string) []string {
			if comp.Workflow() == nil {
				return []string{"empty client"}
			}
			reply, err := comp.Workflow().ActionDoc(ctx, &pb.WorkflowRequest{})
			if err != nil {
				return []string{"error call: " + err.Error()}
			}
			return []string{reply.GetText()}
		},
	},
	{
		Regex: `test`,
		Help:  `Test`,
		Parse: func(ctx context.Context, comp rulebot.IComponent, s string, args []string) []string {
			if comp.Storage() == nil {
				return []string{"empty client"}
			}
			// upload
			f, err := os.Open("./README.md")
			if err != nil {
				return []string{"error: " + err.Error()}
			}

			buf := make([]byte, 1024)
			uc, err := comp.Storage().UploadFile(ctx)
			if err != nil {
				return []string{"error: " + err.Error()}
			}

			err = uc.Send(&pb.FileRequest{Data: &pb.FileRequest_Info{Info: &pb.FileInfo{FileType: "md"}}})
			if err != nil {
				return []string{"error: " + err.Error()}
			}

			for {
				n, err := f.Read(buf)
				if errors.Is(err, io.EOF) {
					break
				}
				if err != nil {
					return []string{"error: " + err.Error()}
				}
				err = uc.Send(&pb.FileRequest{Data: &pb.FileRequest_Chuck{Chuck: buf[:n]}})
				if err != nil {
					return []string{"error: " + err.Error()}
				}
			}

			_, err = uc.CloseAndRecv()
			if err != nil {
				return []string{"error: " + err.Error()}
			}

			return []string{"test done"}
		},
	},
	{
		Regex: `stats`,
		Help:  `Stats Info`,
		Parse: func(ctx context.Context, comp rulebot.IComponent, s string, args []string) []string {
			if comp.Middle() == nil {
				return []string{"empty client"}
			}
			reply, err := comp.Middle().GetStats(ctx, &pb.TextRequest{})
			if err != nil {
				return []string{"error call: " + err.Error()}
			}
			return []string{reply.GetText()}
		},
	},
	{
		Regex: `todo\s+(.*)`,
		Help:  "Todo something",
		Parse: func(ctx context.Context, comp rulebot.IComponent, s string, args []string) []string {
			if comp.Todo() == nil {
				return []string{"empty client"}
			}
			if len(args) != 2 {
				return []string{"error args"}
			}
			reply, err := comp.Todo().CreateTodo(ctx, &pb.TodoRequest{
				Todo: &pb.Todo{Content: args[1]},
			})
			if err != nil {
				return []string{"error call: " + err.Error()}
			}
			if !reply.GetState() {
				return []string{"failed"}
			}
			return []string{"success"}
		},
	},
	{
		Regex: `role`,
		Help:  "Role info",
		Parse: func(ctx context.Context, comp rulebot.IComponent, s string, args []string) []string {
			if comp.Middle() == nil {
				return []string{"empty client"}
			}
			reply, err := comp.Middle().GetRoleImageUrl(ctx, &pb.TextRequest{})
			if err != nil {
				return []string{"error call: " + err.Error()}
			}
			return []string{reply.GetText()}
		},
	},
	{
		Regex: `pinyin\s+(.*)`,
		Help:  "chinese pinyin conversion",
		Parse: func(ctx context.Context, comp rulebot.IComponent, s string, args []string) []string {
			if comp.NLP() == nil {
				return []string{"empty client"}
			}
			if len(args) != 2 {
				return []string{"error args"}
			}
			reply, err := comp.NLP().Pinyin(ctx, &pb.TextRequest{Text: args[1]})
			if err != nil {
				return []string{"error call: " + err.Error()}
			}
			if len(reply.GetText()) <= 0 {
				return []string{"failed"}
			}
			return []string{strings.Join(reply.GetText(), ", ")}
		},
	},
	{
		Regex: `remind\s+(\w+)\s+(\w+)`,
		Help:  `Remind something`,
		Parse: func(ctx context.Context, comp rulebot.IComponent, s string, args []string) []string {
			if len(args) != 3 {
				return []string{"error args"}
			}

			arg1 := args[1]
			arg2 := args[2]
			fmt.Println(arg1, arg2) // todo remind message

			return []string{}
		},
	},
	{
		Regex: `del\s+(\d+)`,
		Help:  `Delete message`,
		Parse: func(ctx context.Context, comp rulebot.IComponent, s string, args []string) []string {
			if len(args) != 2 {
				return []string{"error args"}
			}

			idStr := args[1]
			id, err := strconv.Atoi(idStr)
			if err != nil {
				return []string{"error call: " + err.Error()}
			}

			fmt.Println(id) // todo delete message

			return []string{}
		},
	},
}

var Options = []rulebot.Option{
	rulebot.RegisterRuleset(New(rules)),
}
