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

func (o *OpenAI) Init(config util.LLMConfig) error {
	o.config = config
	return nil
}

func (o *OpenAI) GenerateCommand(prompt string) (string, error) {
	response, err := o.generate(prompt)
	if err != nil {
		return "", err
	}
	return response.Choices[0].Message.Content, nil
}

func (o *OpenAI) generate(prompt string) (openaiResponse, error) {
	systemPrompt, err := util.GetSystemPrompt()
	if err != nil {
		return openaiResponse{}, err
	}
	payload := requestPayload{
		Model: o.config.Model,
		Messages: []struct {
			Role    string "json:\"role\""
			Content string "json:\"content\""
		}{
			{Role: "system", Content: systemPrompt},
			{Role: "user", Content: prompt},
		},
	}
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return openaiResponse{}, err
	}

	req, err := http.NewRequest("POST", o.config.BaseURL, bytes.NewBuffer(payloadBytes))
	if err != nil {
		return openaiResponse{}, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+o.config.ApiKey)
	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return openaiResponse{}, err
	}
	defer resp.Body.Close()

	// Read and print the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return openaiResponse{}, err
	}

	var response openaiResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println("Error unmarshalling response body:", err)
	}
	return response, nil
}
