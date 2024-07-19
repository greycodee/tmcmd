package util

import "fmt"

// type Prompt struct {
// 	SystemPrompt string
// 	UserPrompt   string
// }

func GetPrompt(userPrompt string) string {
	sysInfoTemplate := `## OS Info
OS: %s
OS Name: %s
Arch: %s
`
	sysInfo := GetSysInfo()
	sysInfoStr := fmt.Sprintf(sysInfoTemplate, sysInfo.OS, sysInfo.OSName, sysInfo.Arch)
	prePrompt := sysInfoStr + `
## Task
Please generate me a  terminal command directly for the following task:
%s

## Note
Please note: 
- Please recommend a command that matches the current system
- Do not include any symbols that are not related to the command
- Do not include any Markdown syntax!
- Do not include any explanations or descriptions, just the command itself. 
- If the task involves a file or directory path, use placeholders (e.g., "/path/to/file"). 
- If necessary, specify the required options or parameters.
`
	return fmt.Sprintf(prePrompt, userPrompt)
}
