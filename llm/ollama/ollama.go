package ollama

import (
	"bytes"
	"encoding/json"
	"fmt"
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

func (o *Ollama) Init(config util.LLMConfig) {
	o.config = config
}

func (o *Ollama) GenerateCommand(prompt string) string {
	return o.generate(prompt)
}

func (o *Ollama) generate(prompt string) string {
	return o.requestLocalOllamaAPI(prompt)
}

func (o *Ollama) requestLocalOllamaAPI(prompt string) string {
	payload := payload{
		Model: o.config.Model,
		Messages: []struct {
			Role    string `json:"role"`
			Content string `json:"content"`
		}{
			{
				Role:    "system",
				Content: util.GetSystemPrompt(),
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
		fmt.Println("Error marshalling request body:", err)
		return ""
	}

	resp, err := http.Post(o.config.BaseURL, "application/json", bytes.NewBuffer(jsonBody))
	if err != nil {
		fmt.Println("Error making request to local ollama server API:", err)
		return ""
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return ""
	}

	// Unmarshal the response body into the responseInfo struct
	var respInfo response
	err = json.Unmarshal(body, &respInfo)
	if err != nil {
		fmt.Println("Error unmarshalling response body:", err)
		return ""
	}
	// Return the response
	return respInfo.Message.Content
}
