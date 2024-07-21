package google

import (
	"context"
	"fmt"

	"github.com/google/generative-ai-go/genai"
	"github.com/greycodee/tmcmd/util"
	"google.golang.org/api/option"
)

type Gemini struct {
	config util.LLMConfig
}

func (g *Gemini) Init(config util.LLMConfig) error {
	g.config = config
	return nil
}

func (g *Gemini) GenerateCommand(userPrompt string) (string, error) {
	systemPrompt, err := util.GetSystemPrompt()
	if err != nil {
		return "", err
	}
	return g.generate(systemPrompt + userPrompt)
}

func (g *Gemini) generate(prompt string) (string, error) {
	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(g.config.ApiKey))
	if err != nil {
		return "", err
	}
	defer client.Close()

	model := client.GenerativeModel(g.config.Model)

	resp, err := model.GenerateContent(ctx, genai.Text(prompt))
	if err != nil {
		return "", err
	}
	result := fmt.Sprint(resp.Candidates[0].Content.Parts[0])
	return result, nil
}
