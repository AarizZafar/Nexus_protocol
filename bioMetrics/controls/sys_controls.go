package controls

import (
	"fmt"
	"os/exec"
	"runtime"
	"strings"

	"github.com/AarizZafar/Nexus_protocol.git/bioMetrics/commands"
)

var currOS = runtime.GOOS
var executeCmd = map[string][]string{}

func init() {
	switch currOS {
	case "linux":
		executeCmd = commands.Linux
	case "darwin":
		executeCmd = commands.Darwin
	case "windows":
		executeCmd = commands.Windows
	}
}

func extractssid(output string) string {
	var lines = strings.Split(output, "\n")

	var SSID = strings.Split(strings.TrimSpace(lines[8]), ":")[1]
	SSID = strings.TrimSpace(SSID)

	if SSID != "" {
		return SSID
	}
	return ""
}

func GetNetssid() string {
	return getSysInfo(
		executeCmd["SSID"],
		extractssid,
	)
}

func safeCommandOutput(commands string, args ...string) string {
	cmd := exec.Command(commands, args...) 	// creates a command that will run in the OS (like running in the terminal)
											// returns a *exec.Cmd (pointer) to exec.Cmd struct -> represents the command to be run in the OS
	output, err := cmd.Output() 			// returns -> []byte - result of the command, error
	if err != nil {
		fmt.Printf("Error executing command: %v\n", err)
		return ""
	}
	return strings.TrimSpace(string(output)) // converts the output -> string, removes any leading or trailing spaces,
}

func getSysInfo(command []string, formatStr func(string) string) string {
	var output string
	output = safeCommandOutput(command[0], command[1:]...)
	if formatStr != nil {
		output = formatStr(output)
	}
	return output
}

// ********************* MAC ********************
func GetMACAdd() string {
	var output = safeCommandOutput(executeCmd["MACAddress"][0], executeCmd["MACAddress"][1:]...)
	
	lines := strings.Split(output, "\n") // splitting the raw output string into a slice of string based on spaces
	parts := strings.Fields(lines[2])    // return type is a slice of string
	if len(parts) > 0 {
		return parts[0]
	}
	return ""
}

// *********** SYSTEM SERIAL NUMBER *************
func extractSystemSerialNumber(output string) string{
	lines := strings.Split(output, "\n")
	var trimmed_data []string

	if len(lines) > 1 && strings.TrimSpace(lines[1]) != "" {
		lines[1] = strings.TrimSpace(lines[1])
		trimmed_data = append(trimmed_data, lines[1])
		return trimmed_data[0]
	} else if len(lines) <= 1 {
		fmt.Println("\033[43;30m ⚠️ Warning: Output does not contain enough lines. ⚠️  \033[0m")
	} else {
		fmt.Println("\033[43;30m ⚠️ Warning: Output line is empty.                 ⚠️  \033[0m")
	}
	return ""
}

func GetSystemSerialNumber() string {
	return getSysInfo(
		executeCmd["SystemSerialNumber"],
		extractSystemSerialNumber,
	)
}

// ************* SYSTEM UUID *******************
func extractSystemUUID(output string) string {
	lines := strings.Split(output, "\n")
	var trimmed_data []string 

	if len(lines) > 1 && strings.TrimSpace(lines[1]) != "" {
		lines[1] = strings.TrimSpace(lines[1])
		trimmed_data = append(trimmed_data, lines[1])
		return trimmed_data[0]
	} else if len(lines) <= 1 {
		fmt.Println("\033[43;30m ⚠️ Warning: Output does not contain enough lines. ⚠️  \033[0m")
	} else {
		fmt.Println("\033[43;30m ⚠️ Warning: Output line is empty.                 ⚠️  \033[0m")
	}
	return ""
}

func GetSystemUUID() string {
	return getSysInfo(
		executeCmd["UUID"],
		extractSystemUUID,
	)
}