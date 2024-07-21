package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/greycodee/tmcmd/llm"
	"github.com/greycodee/tmcmd/llm/google"
	"github.com/greycodee/tmcmd/llm/ollama"
	"github.com/greycodee/tmcmd/llm/openai"
	"github.com/greycodee/tmcmd/util"
)

func main() {
	// Accept command line arguments
	userPrompt := os.Args[1:]
	// prompt := util.GetPrompt(strings.Join(userPrompt, ""))
	if len(userPrompt) == 0 {
		userPrompt = append(userPrompt, "查询当前系统发行版本")
	}
	// Get the configuration
	config, err := util.GetConfig()
	if err != nil {
		fmt.Println("Error getting config:", err)
		return
	}
	var llm llm.LLMBaseInterface
	switch config.DefaultProvider {
	case "ollama":
		// Initialize ollama
		llm = new(ollama.Ollama)
	case "google":
		// Initialize google
		llm = new(google.Gemini)
	case "openai":
		// Initialize openai
		llm = new(openai.OpenAI)
	default:
		fmt.Println("Default provider not supported")
	}
	llm.Init(config.LLMProvider[config.DefaultProvider])
	command := llm.GenerateCommand(strings.Join(userPrompt, ""))
	fmt.Printf("\033[42mRecommended command\033[0m\r\n\033[32m%s\033[0m\n", command)
}
