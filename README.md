# tmcmd (tell me commands)

tmcmd is a powerful terminal tool designed to enhance your productivity by providing AI-powered command recommendations. With a simple and intuitive interface, you can easily obtain suggestions for commands based on your input, streamlining your workflow and saving valuable time.

## Features

- **Flexible Configuration**: Customize the tool to suit your needs by modifying the default command provider directly from the configuration file.
- **Provider Selection**: Choose from a variety of language model providers for each request, giving you the flexibility to select the best option for your specific needs.
- **Interactive Prompts**: Simply enter your prompt, and let the AI recommend the most relevant commands to accomplish your tasks efficiently.

## Options

- `-default <string>`: Modify the default provider specified in the configuration file.
- `-p <string>`: Set the LLM (Large Language Model) provider for the current request. This option is not mandatory.
- `-q <string>`: Enter your prompt to receive command recommendations.

## Supported LLM Providers

The tool supports a wide range of LLM providers, including:

- Ollama
- Google Gemini
- All platforms that adhere to the OpenAI interface specification.

This ensures that you have access to a diverse set of AI technologies to get the most accurate and relevant command recommendations.

## Config File Path

- **Windows**: `$HOME/AppData/Local/tmcmd/config.toml`
- **MacOS**: `$HOME/.config/tmcmd/config.toml`
- **Linux**: `$HOME/.config/tmcmd/config.toml`

This is a sample configuration [config_example.toml](./config_example.toml)

## Manual Installation

```base
go install github.com/greycodee/tmcmd@latest
```

Or you can compile from source

```bash
git clone https://github.com/greycodee/tmcmd.git
cd tmcmd
go build .
```
