package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/greycodee/tmcmd/llm/ollama"
	"github.com/greycodee/tmcmd/util"
)

func main() {
	// Accept command line arguments
	userPrompt := os.Args[1:]
	prompt := util.GetPrompt(strings.Join(userPrompt, ""))
	// 初始化ollama
	llm := new(ollama.Ollama)
	command := llm.GenerateCommand(prompt)
	fmt.Printf("Recommended command: %s\n", command)
}
