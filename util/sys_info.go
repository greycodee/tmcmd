package util

import (
	"bufio"
	"bytes"
	"os"
	"runtime"
	"strings"
)

type SysInfo struct {
	OS     string
	OSName string
	Arch   string
}

func GetSysInfo() *SysInfo {
	sysInfo := &SysInfo{
		OS:     runtime.GOOS,
		OSName: getOSName(),
		Arch:   runtime.GOARCH,
	}
	return sysInfo
}

func getOSName() string {
	switch runtime.GOOS {
	case "darwin":
		return "macOS"
	case "linux":
		return getLinuxDistro()
	case "windows":
		return "Windows"
	default:
		return "Unknown"
	}
}

func getLinuxDistro() string {
	filePath := "/etc/os-release"
	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		return "Unknown"
	}

	scanner := bufio.NewScanner(bytes.NewReader(fileContent))
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "PRETTY_NAME=") {
			return strings.TrimPrefix(line, "PRETTY_NAME=")
		}
	}

	return "Unknown"
}

// var cmd *exec.Cmd
// switch runtime.GOOS {
// case "linux":
// 	// 在大多数 Linux 发行版上，lsb_release 命令可以提供发行版信息
// 	cmd = exec.Command("lsb_release", "-a")
// case "darwin":
// 	// 在 macOS 上，sw_vers 命令可以提供系统版本信息
// 	cmd = exec.Command("sw_vers")
// case "windows":
// 	// 在 Windows 上，systeminfo 命令可以提供系统版本信息
// 	cmd = exec.Command("systeminfo")
// default:
// 	fmt.Println("Unsupported operating system")
// 	return
// }

// out, err := cmd.Output()
// if err != nil {
// 	fmt.Println("Error getting system information:", err)
// 	return
// }

// // 根据不同的操作系统解析输出
// switch runtime.GOOS {
// case "linux":
// 	for _, line := range strings.Split(string(out), "\n") {
// 		if strings.HasPrefix(line, "Description:") {
// 			fmt.Println(strings.TrimSpace(strings.TrimPrefix(line, "Description:")))
// 		}
// 	}
// case "darwin":
// 	for _, line := range strings.Split(string(out), "\n") {
// 		if strings.HasPrefix(line, "ProductVersion:") {
// 			fmt.Println(strings.TrimSpace(strings.TrimPrefix(line, "ProductVersion:")))
// 		}
// 		if strings.HasPrefix(line, "ProductName:") {
// 			fmt.Println(strings.TrimSpace(strings.TrimPrefix(line, "ProductName:")))
// 		}
// 	}

// case "windows":
// 	for _, line := range strings.Split(string(out), "\r\n") {
// 		if strings.HasPrefix(line, "OS Name:") {
// 			fmt.Println(strings.TrimSpace(strings.TrimPrefix(line, "OS Name:")))
// 			return
// 		}
// 	}
// }
