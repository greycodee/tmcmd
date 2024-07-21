package util

import (
	"bufio"
	"bytes"
	"errors"
	"os"
	"runtime"
	"strings"
)

type SysInfo struct {
	OS     string
	OSName string
	Arch   string
}

func GetSysInfo() (*SysInfo, error) {
	osName, err := getOSName()
	if err != nil {
		return nil, err
	}

	sysInfo := &SysInfo{
		OS:     runtime.GOOS,
		OSName: osName,
		Arch:   runtime.GOARCH,
	}
	return sysInfo, nil
}

func getOSName() (string, error) {
	switch runtime.GOOS {
	case "darwin":
		return "macOS", nil
	case "linux":
		return getLinuxDistro()
	case "windows":
		return "Windows", nil
	default:
		return "", errors.New("unsupported OS")
	}
}

func getLinuxDistro() (string, error) {
	filePath := "/etc/os-release"
	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	scanner := bufio.NewScanner(bytes.NewReader(fileContent))
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "PRETTY_NAME=") {
			return strings.TrimPrefix(line, "PRETTY_NAME="), nil
		}
	}

	return "Unknown", nil
}
