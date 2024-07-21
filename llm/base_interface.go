package llm

import "github.com/greycodee/tmcmd/util"

type LLMBaseInterface interface {
	GenerateCommand(prompt string) string
	Init(config util.LLMConfig)
}
