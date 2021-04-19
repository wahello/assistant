package stage

import (
	"encoding/json"
	"github.com/tsundata/assistant/internal/app/cron/pipeline/result"
	"github.com/tsundata/assistant/internal/pkg/rulebot"
	"github.com/tsundata/assistant/internal/pkg/utils"
)

func URL(_ *rulebot.RuleBot, r result.Result) result.Result {
	if r.Kind == result.Url {
		j, err := json.Marshal(r.Content)
		if err != nil {
			return result.ErrorResult(err)
		}
		return result.MessageResult(utils.ByteToString(j))
	}
	return result.EmptyResult()
}