package bioMetrics

import (
	"fmt"
	"os/exec"   // used to run external commands or progs from withing a Go program and capture their output
	"runtime"   // provides information obout the go runtime system
	"strings"
)

var currOS = runtime.GOOS

func safeCommandOutput(commands string, args ...string) string {
	cmd := exec.Command(commands, args...) // executing the command in the os
	output, err := cmd.Output()
	if err != nil {
		return ""
	}
	return strings.TrimSpace(string(output))
}

func getMACAdd() string {
	var output string

	switch currOS {
	case "linux", "darwin":      // linux or macOS
			output = safeCommandOutput("ifconfid")
	case "windows":              // windows
		    output = safeCommandOutput("getmac")
	default:
			return ""
	}

	if output == "" {
		return ""
	}

	lines := strings.Split(output, "\n")
	for _, line := range lines {
		if strings.Contains(line, "ether") || strings.Contains(line, "-") {
			parts := strings.Fields(line)
			if len(parts) > 0 {
				return parts[0]   // first part is usually the MAC address
			}
		}
	}
	return ""
}

func ExtSysBioMetrix() {
	macAdd := getMACAdd()

	fmt.Println("The mac addres is : ", macAdd)
}