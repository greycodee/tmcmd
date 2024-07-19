package google

type Gemini struct {
}

func (g *Gemini) GenerateCommand(userPrompt string) string {
	return g.generate()
}

func (g *Gemini) generate() string {
	return "echo Hello, Gemini!"
}
