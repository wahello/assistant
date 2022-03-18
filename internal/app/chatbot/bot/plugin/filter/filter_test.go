package filter

import (
	"context"
	"github.com/stretchr/testify/assert"
	end2 "github.com/tsundata/assistant/internal/app/chatbot/bot/plugin/end"
	"github.com/tsundata/assistant/internal/pkg/robot/bot"
	"testing"
)

func TestFilter(t *testing.T) {
	p := Filter{
		Next: end2.End{},
	}
	input := "test"
	out, err := p.Run(context.Background(), bot.MockController(), input)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, input, out)
}