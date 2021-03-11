package rules

import (
	"github.com/tsundata/assistant/internal/app/cron/agent"
	"github.com/tsundata/assistant/internal/pkg/rulebot"
)

var rules = []Rule{
	{
		Name: "pocket",
		When: "* * * * *",
		Action: func(b *rulebot.RuleBot) []string {
			return agent.FetchPocket(b)
		},
	},
	{
		Name: "github_starred",
		When: "* * * * *",
		Action: func(b *rulebot.RuleBot) []string {
			return agent.FetchGithubStarred(b)
		},
	},
	{
		Name: "backup",
		When: "0 0 * * *",
		Action: func(b *rulebot.RuleBot) []string {
			return agent.Backup(b)
		},
	},
}

var Options = []rulebot.Option{
	rulebot.RegisterRuleset(New(rules)),
}
