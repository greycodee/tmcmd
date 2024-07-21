package google

import (
	"context"
	"fmt"
	"log"

	"github.com/google/generative-ai-go/genai"
	"github.com/greycodee/tmcmd/util"
	"google.golang.org/api/option"
)

type Gemini struct {
	config util.LLMConfig
}

func (g *Gemini) Init(config util.LLMConfig) {
	g.config = config
}

func (g *Gemini) GenerateCommand(userPrompt string) string {
	return g.generate(util.GetSystemPrompt() + userPrompt)
}

func (g *Gemini) generate(prompt string) string {
	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(g.config.ApiKey))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	model := client.GenerativeModel(g.config.Model)

	resp, err := model.GenerateContent(ctx, genai.Text(prompt))
	if err != nil {
		log.Fatal(err)
	}
	return fmt.Sprint(resp.Candidates[0].Content.Parts[0])
}
