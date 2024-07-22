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

func getConfigFilePath() (string, error) {
	configPath := os.Getenv("TMCMD_CONFIG_PATH")
	if configPath != "" {
		return configPath, nil
	}
	configPath, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	switch runtime.GOOS {
	case "windows":
		configPath += "/AppData/Local/tmcmd"
	case "darwin":
		configPath += "/.config/tmcmd"
	case "linux":
		configPath += "/.config/tmcmd"
	default:
		return "", fmt.Errorf("unsupported OS: %s", runtime.GOOS)
	}
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		err = os.MkdirAll(configPath, os.ModePerm)
		if err != nil {
			return "", err
		}
	}
	return configPath + "/config.toml", nil
}

func GetConfig() (*Config, error) {
	configFile, err := getConfigFilePath()
	if err != nil {
		return nil, err
	}
	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		_, err := os.Create(configFile)
		if err != nil {
			return nil, err
		}
		defaultConfig := Config{
			DefaultProvider: "ollama",
			LLMProvider: map[string]LLMConfig{
				"google": {
					BaseURL: "No need to fill in",
					Model:   "gemini-1.5-pro",
					ApiKey:  "YOUR_API_KEY",
				},
				"ollama": {
					BaseURL: "http://localhost:11434/api/chat",
					Model:   "llama3",
				},
				"openai": {
					BaseURL: "https://api.deepinfra.com/v1/openai/chat/completions",
					Model:   "google/gemma-2-27b-it",
					ApiKey:  "YOUR_API_KEY",
				},
			},
		}
		writeConfig(configFile, &defaultConfig)
	}
	return ParseConfig(configFile)
}

// ParseConfig reads and parses the TOML configuration file
func ParseConfig(filePath string) (*Config, error) {
	var config Config
	if _, err := toml.DecodeFile(filePath, &config); err != nil {
		return nil, err
	}
	return &config, nil
}

func SetDefaultProvider(provider string) error {
	if !IsSupportedProvider(provider) {
		return fmt.Errorf("unsupported provider: %s", provider)
	}
	config, err := GetConfig()
	if err != nil {
		return err
	}
	config.DefaultProvider = provider
	filePath, err := getConfigFilePath()
	if err != nil {
		return err
	}
	return writeConfig(filePath, config)
}

func writeConfig(filePath string, config *Config) error {
	configBytes, err := toml.Marshal(config)
	if err != nil {
		return err
	}
	return os.WriteFile(filePath, configBytes, os.ModePerm)
}

func IsSupportedProvider(provider string) bool {
	switch provider {
	case "openai", "google", "ollama":
		return true
	default:
		return false
	}
}
