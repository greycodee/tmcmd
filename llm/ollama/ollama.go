package ollama

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/greycodee/tmcmd/util"
)

type Ollama struct {
	config util.LLMConfig
}

type payload struct {
	Model    string `json:"model"`
	Messages []struct {
		Role    string `json:"role"`
		Content string `json:"content"`
	} `json:"messages"`
	Stream bool `json:"stream"`
}

type response struct {
	Model     string    `json:"model"`
	CreatedAt time.Time `json:"created_at"`
	Message   struct {
		Role    string `json:"role"`
		Content string `json:"content"`
	} `json:"message"`
	Done               bool  `json:"done"`
	TotalDuration      int64 `json:"total_duration"`
	LoadDuration       int   `json:"load_duration"`
	PromptEvalCount    int   `json:"prompt_eval_count"`
	PromptEvalDuration int   `json:"prompt_eval_duration"`
	EvalCount          int   `json:"eval_count"`
	EvalDuration       int64 `json:"eval_duration"`
}

func (o *Ollama) Init(config util.LLMConfig) error {
	o.config = config
	return nil
}

func (o *Ollama) GenerateCommand(prompt string) (string, error) {
	return o.generate(prompt)
}

func (o *Ollama) generate(prompt string) (string, error) {
	return o.requestLocalOllamaAPI(prompt)
}

func (o *Ollama) requestLocalOllamaAPI(prompt string) (string, error) {
	systemPrompt, err := util.GetSystemPrompt()
	if err != nil {
		return "", err
	}
	payload := payload{
		Model: o.config.Model,
		Messages: []struct {
			Role    string `json:"role"`
			Content string `json:"content"`
		}{
			{
				Role:    "system",
				Content: systemPrompt,
			},
			{
				Role:    "user",
				Content: prompt,
			},
		},
		Stream: false,
	}

	jsonBody, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}

	resp, err := http.Post(o.config.BaseURL, "application/json", bytes.NewBuffer(jsonBody))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	// Unmarshal the response body into the responseInfo struct
	var respInfo response
	err = json.Unmarshal(body, &respInfo)
	if err != nil {
		return "", err
	}
	// Return the response
	return respInfo.Message.Content, nil
}
