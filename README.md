# tmcmd (tell me commands)

tmcmd is a powerful terminal tool designed to enhance your productivity by providing AI-powered command recommendations. With a simple and intuitive interface, you can easily obtain suggestions for commands based on your input, streamlining your workflow and saving valuable time.

> Note: The accuracy of the recommended commands is related to the strength of the model, the stronger the model, the more accurate the recommended commands are.

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

## Usage

```bash
tmcmd -q "View 80-port occupancy status"

Recommended command
netstat -a -n -p tcp | grep :80
```

```bash
tmcmd -q "Start the nginx service on docker and expose port 80."

Recommended command
docker run -d -p 80:80 ngin
```

Use the `-p ollama` option to select ollama to process the current request.

```bash
tmcmd -q "Start the nginx service on docker and expose port 80." -p ollama

Recommended command
docker run -d -p 80:80 ngin
```
