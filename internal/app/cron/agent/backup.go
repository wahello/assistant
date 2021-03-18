package agent

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/rulebot"
	"github.com/tsundata/assistant/internal/pkg/vendors/dropbox"
	"time"
)

func Backup(b *rulebot.RuleBot) []string {
	ctx := context.Background()
	app, err := b.MidClient.GetAvailableApp(ctx, &pb.TextRequest{Text: dropbox.ID})
	if err != nil {
		return []string{err.Error()}
	}
	accessToken := app.GetToken()
	if accessToken == "" {
		return []string{"backup: dropbox access token is empty"}
	}

	// messages
	messagesReply, err := b.MsgClient.List(ctx, &pb.MessageRequest{})
	if err != nil {
		return []string{err.Error()}
	}

	// apps
	appsReply, err := b.MidClient.GetApps(ctx, &pb.TextRequest{})
	if err != nil {
		return []string{err.Error()}
	}

	// credentials
	credentialsReply, err := b.MidClient.GetCredentials(ctx, &pb.TextRequest{})
	if err != nil {
		return []string{err.Error()}
	}

	data := map[string]interface{}{
		"message":     messagesReply.Messages,
		"apps":        appsReply.Apps,
		"credentials": credentialsReply.Credentials,
	}
	d, err := json.Marshal(data)
	if err != nil {
		return []string{err.Error()}
	}

	// upload
	client := dropbox.NewDropbox("", "", "", accessToken)
	path := fmt.Sprintf("/backup/assistant_%s.json", time.Now().Format("2006-01-02_15:04:05"))
	err = client.Upload(path, bytes.NewReader(d))
	if err != nil {
		return []string{err.Error()}
	}

	return []string{fmt.Sprintf("backed up in %s", time.Now())}
}