package util

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseConfigValidFile(t *testing.T) {
	content := `
default-provider = "example"
[llm-provider."provider1"]
base_url = "http://example.com"
mode = "test"
`
	// Use os.CreateTemp instead of ioutil.TempFile
	tmpFile, err := os.CreateTemp("", "config-*.toml")
	assert.NoError(t, err)
	defer os.Remove(tmpFile.Name()) // Clean up the file after the test

	_, err = tmpFile.WriteString(content)
	assert.NoError(t, err)
	tmpFile.Close()

	config, err := ParseConfig(tmpFile.Name())
	assert.NoError(t, err)
	assert.NotNil(t, config)
	assert.Equal(t, "example", config.DefaultProvider)
	assert.Equal(t, "http://example.com", config.LLMProvider["provider1"].BaseURL)
	assert.Equal(t, "test", config.LLMProvider["provider1"].Model)
}

func TestParseConfigInvalidFile(t *testing.T) {
	content := `
default-provider = "example"
[llm-provider.provider1]
base_url = "http://example.com"
mode = "test"
invalid_field = "should_fail"
`
	// Use os.CreateTemp instead of ioutil.TempFile
	tmpFile, err := os.CreateTemp("", "config-*.toml")
	assert.NoError(t, err)
	defer os.Remove(tmpFile.Name()) // Clean up the file after the test

	_, err = tmpFile.WriteString(content)
	assert.NoError(t, err)
	tmpFile.Close()

	config, err := ParseConfig(tmpFile.Name())
	assert.Nil(t, err)
	assert.NotNil(t, config)
}

func TestParseConfigNonExistentFile(t *testing.T) {
	config, err := ParseConfig("nonexistent.toml")
	assert.Error(t, err)
	assert.Nil(t, config)
}

func TestParseConfigFilePath(t *testing.T) {
	path := "../config.toml"
	config, err := ParseConfig(path)
	assert.NoError(t, err)
	assert.NotNil(t, config)

}
