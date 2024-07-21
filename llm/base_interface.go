package llm

type LLMBaseInterface interface {
	GenerateCommand(prompt string) string
}
