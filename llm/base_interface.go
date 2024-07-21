package llm

import "github.com/greycodee/tmcmd/util"

type LLMBaseInterface interface {
	GenerateCommand(prompt string) (string, error)
	Init(config util.LLMConfig) error
}
