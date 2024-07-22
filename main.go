package main

import (
	"errors"
	"flag"
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

var (
	defaultProvider = flag.String("default", "", "modify the default-provider of the configuration file")
	requestProvider = flag.String("p", "", "set the llm provider for this request")
	query           = flag.String("q", "", "enter your prompt")
)

func init() {
	flag.Parse()
	if *defaultProvider != "" {
		err := util.SetDefaultProvider(*defaultProvider)
		if err != nil {
			logError(err)
		}
		logInfo(fmt.Sprintf("set default provider: %s\n", *defaultProvider))
	}

	if *query == "" {
		logError(errors.New("please use -q to enter your prompt"))
	}
}

func main() {
	s := spinner.New(spinner.CharSets[34], 100*time.Millisecond)
	s.Start()
	s.Color("green")
	userPrompt := os.Args[1:]
	if len(userPrompt) == 0 {
		userPrompt = append(userPrompt, "how to query current information")
	}
	config, err := util.GetConfig()
	if err != nil {
		s.Stop()
		logError(err)
		return
	}
	if *requestProvider != "" {
		if !util.IsSupportedProvider(*requestProvider) {
			s.Stop()
			logError(errors.New("unsupported provider: " + *requestProvider))
		}
		config.DefaultProvider = *requestProvider
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

func logInfo(info string) {
	log.Fatalf("\033[32m%s\033[0m\n", info)
}
