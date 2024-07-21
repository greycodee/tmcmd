package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/briandowns/spinner"
	"github.com/greycodee/tmcmd/llm"
	"github.com/greycodee/tmcmd/llm/google"
	"github.com/greycodee/tmcmd/llm/ollama"
	"github.com/greycodee/tmcmd/llm/openai"
	"github.com/greycodee/tmcmd/util"
)

func main() {
	s := spinner.New(spinner.CharSets[34], 100*time.Millisecond)
	s.Start()
	s.Color("green")
	// Accept command line arguments
	userPrompt := os.Args[1:]
	// prompt := util.GetPrompt(strings.Join(userPrompt, ""))
	if len(userPrompt) == 0 {
		userPrompt = append(userPrompt, "how to query current information")
	}
	// Get the configuration
	config, err := util.GetConfig()
	if err != nil {
		s.Stop()
		logError(err)
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
		s.Stop()
		logError(errors.New("default provider not supported"))
	}
	llm.Init(config.LLMProvider[config.DefaultProvider])
	command, err := llm.GenerateCommand(strings.Join(userPrompt, ""))
	if err != nil {
		s.Stop()
		logError(err)
	}
	s.Stop()
	fmt.Printf("\033[42mRecommended command\033[0m\r\n\033[32m%s\033[0m\n", command)
}

func logError(err error) {
	if err != nil {
		log.Fatalf("\033[31mError: %v\033[0m\n", err)
	}
}
