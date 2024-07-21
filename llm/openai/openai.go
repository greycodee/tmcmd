package openai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/greycodee/tmcmd/util"
)

type OpenAI struct {
	config util.LLMConfig
}

type requestPayload struct {
	Model    string `json:"model"`
	Messages []struct {
		Role    string `json:"role"`
		Content string `json:"content"`
	} `json:"messages"`
}

type openaiResponse struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Created int    `json:"created"`
	Model   string `json:"model"`
	Choices []struct {
		Index   int `json:"index"`
		Message struct {
			Role    string `json:"role"`
			Content string `json:"content"`
		} `json:"message"`
		FinishReason string `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int     `json:"prompt_tokens"`
		CompletionTokens int     `json:"completion_tokens"`
		TotalTokens      int     `json:"total_tokens"`
		EstimatedCost    float64 `json:"estimated_cost"`
	} `json:"usage"`
}

func (o *OpenAI) Init(config util.LLMConfig) {
	o.config = config
}

func (o *OpenAI) GenerateCommand(prompt string) string {
	return o.generate(prompt).Choices[0].Message.Content
}

func (o *OpenAI) generate(prompt string) openaiResponse {
	payload := requestPayload{
		Model: o.config.Model,
		Messages: []struct {
			Role    string "json:\"role\""
			Content string "json:\"content\""
		}{
			{Role: "system", Content: util.GetSystemPrompt()},
			{Role: "user", Content: prompt},
		},
	}
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("Error marshalling payload:", err)
	}

	req, err := http.NewRequest("POST", o.config.BaseURL, bytes.NewBuffer(payloadBytes))
	if err != nil {
		fmt.Println("Error creating request:", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+o.config.ApiKey)
	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
	}
	defer resp.Body.Close()

	// Read and print the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
	}

	var response openaiResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println("Error unmarshalling response body:", err)
	}
	return response
}
