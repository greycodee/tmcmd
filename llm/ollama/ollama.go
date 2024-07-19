package ollama

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Ollama struct {
	response responseInfo
}

type responseInfo struct {
	TotalDuration      int64  `json:"total_duration"`
	LoadDuration       int64  `json:"load_duration"`
	PromptEvalCount    int    `json:"prompt_eval_count"`
	PromptEvalDuration int64  `json:"prompt_eval_duration"`
	EvalCount          int    `json:"eval_count"`
	EvalDuration       int64  `json:"eval_duration"`
	Context            []int  `json:"context"`
	Response           string `json:"response"`
}

func (o *Ollama) GenerateCommand(prompt string) string {
	return o.generate(prompt)
}

func (o *Ollama) generate(prompt string) string {
	return o.requestLocalOllamaAPI(prompt)
}

func (o *Ollama) requestLocalOllamaAPI(prompt string) string {
	requestBody := map[string]interface{}{
		"model":  "llama3",
		"prompt": prompt,
		"stream": false,
	}
	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		fmt.Println("Error marshalling request body:", err)
		return ""
	}

	resp, err := http.Post("http://localhost:11434/api/generate", "application/json", bytes.NewBuffer(jsonBody))
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
	o.response = responseInfo{}
	err = json.Unmarshal(body, &o.response)
	if err != nil {
		fmt.Println("Error unmarshalling response body:", err)
		return ""
	}
	// Return the response
	return o.response.Response
}
