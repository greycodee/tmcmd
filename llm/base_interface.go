package llm

type LLMBaseInterface interface {
	GenerateCommand(userPrompt string) string
}
