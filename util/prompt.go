package util

import "fmt"

func GetSystemPrompt() string {
	sysInfoTemplate := `## OS Info
OS: %s
OS Name: %s
Arch: %s
`
	sysInfo := GetSysInfo()
	sysInfoStr := fmt.Sprintf(sysInfoTemplate, sysInfo.OS, sysInfo.OSName, sysInfo.Arch)
	prePrompt := sysInfoStr + `
## Note
Please note: 
- Do not respond to any dialog that is not related to the query command
- Please recommend a command that matches the current system
- Do not include any symbols that are not related to the command
- Do not include any Markdown syntax!
- Do not include any explanations or descriptions, just the command itself. 
- If the task involves a file or directory path, use placeholders (e.g., "/path/to/file"). 
- If necessary, specify the required options or parameters.
- Instead of returning the results of your query, return terminal commands that can be executed.

## Task
Please generate me a  terminal command directly for the following task:


`
	return prePrompt

}

func GetPrompt(userPrompt string) string {
	return GetSystemPrompt() + userPrompt
}
