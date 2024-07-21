package util

import (
	"fmt"
	"os"
	"runtime"

	"github.com/BurntSushi/toml"
)

// Config represents the top-level configuration structure
type Config struct {
	DefaultProvider string               `toml:"default-provider"`
	LLMProvider     map[string]LLMConfig `toml:"llm-provider"`
}

// LLMConfig represents the configuration for an LLM provider
type LLMConfig struct {
	BaseURL string `toml:"base_url"`
	Model   string `toml:"model"`
	ApiKey  string `toml:"api_key"`
}

func GetConfig() (*Config, error) {
	configPath := os.Getenv("TMCMD_CONFIG_PATH")
	if configPath != "" {
		return ParseConfig(configPath)
	}
	configPath, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}
	switch runtime.GOOS {
	case "windows":
		configPath += "/AppData/Local/tmcmd"
	case "darwin":
		configPath += "/.config/tmcmd"
	case "linux":
		configPath += "/.config/tmcmd"
	default:
		return nil, fmt.Errorf("unsupported OS: %s", runtime.GOOS)
	}
	configFileName := "/config.toml"
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		os.MkdirAll(configPath, os.ModePerm)
	}
	if _, err := os.Stat(configPath + configFileName); os.IsNotExist(err) {
		_, err := os.Create(configPath + configFileName)
		if err != nil {
			return nil, err
		}
		defaultConfig := `
default-provider = "ollama"

[llm-provider.google]
model = "gemini-1.5-pro"
api_key = "your-api-key"

[llm-provider.ollama]
base_url = "http://localhost:11434/api/chat"
model = "llama3"

# Supports both OpenAI and Deepinfra
[llm-provider.openai]
base_url = "https://api.deepinfra.com/v1/openai/chat/completions"
model = "google/gemma-2-27b-it"
api_key = "your-api-key"
`

		err = os.WriteFile(configPath+configFileName, []byte(defaultConfig), os.ModePerm)
		if err != nil {
			return nil, err
		}
	}
	return ParseConfig(configPath + configFileName)
}

// ParseConfig reads and parses the TOML configuration file
func ParseConfig(filePath string) (*Config, error) {
	var config Config
	if _, err := toml.DecodeFile(filePath, &config); err != nil {
		return nil, err
	}
	return &config, nil
}
