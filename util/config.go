package util

import (
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
	return ParseConfig("config.toml")
}

// ParseConfig reads and parses the TOML configuration file
func ParseConfig(filePath string) (*Config, error) {
	var config Config
	if _, err := toml.DecodeFile(filePath, &config); err != nil {
		return nil, err
	}
	return &config, nil
}
