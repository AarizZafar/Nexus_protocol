package controls

import (
	"os/exec"
	"runtime"
	"strings"

	"github.com/AarizZafar/Nexus_protocol.git/bioMetrics/commands"
)

var OS = runtime.GOOS
var Net_commands = map[string][]string {}

func init() {
	switch OS {
	case "linux":
		Net_commands = commands.Net_Linux
	case "darwin":
		Net_commands = commands.Net_darwin
	case "windows":
		Net_commands = commands.Net_Windows
	}
}

func safeNetCommandOutput(command string, args ...string) string {
	cmd := exec.Command(command, args...)   // cannot take shell features like the |(pipe)
	output, err := cmd.Output()
	if err != nil {
		return ""
	}
	return strings.TrimSpace(string(output))
}

func getNetInfo(command []string, formatStr func(string) string) string {
	var output string
	output = safeNetCommandOutput(command[0], command[1:]...)
	if formatStr != nil {
		output = formatStr(output)
	}
	return output
}

// ********************* SSID ********************
func extractSSID(output string) string {
	var lines = strings.Split(output, "\n")

	var SSID = strings.Split(strings.TrimSpace(lines[8]), ":")[1]
	SSID = strings.TrimSpace(SSID)

	if SSID != "" {
		return SSID
	}
	return ""
}

func GetNetSSID() string {
	return getNetInfo(
		Net_commands["SSID"],
		extractSSID,
	)
}

// ******************** BSSID ********************
func extractBSSID(output string) string {
	var lines = strings.Split(output, "\n")
	// there will be many : extract the string after the first :, hence start from 2
	var BSSID = strings.SplitN(strings.TrimSpace(lines[9]),":",2)[1]
	BSSID = strings.TrimSpace(BSSID)
	
	if BSSID != "" {
		return BSSID
	}
	return ""
}

func GetNetBSSID() string {
	return getNetInfo(
		Net_commands["BSSID"],
		extractBSSID,
	)
}

// *************** NET INTERFACE MAC *************
func extractNetInterfaceMAC(output string) string {
	lines := strings.Split(output, "\n")
	var trimmed_data []string

	for _, line := range lines {
		trimmed_line := strings.TrimSpace(line)
		data := strings.Fields(trimmed_line)

		trimmed_data = append(trimmed_data, data...)
	}

	if trimmed_data[6] != "" {
		return trimmed_data[6]
	}

	return ""

}

func GetNetInterfaceMac() string {
	return getNetInfo(
		Net_commands["NetworkInterfaceMAC"],
		extractNetInterfaceMAC,
	)
}

// ***************** SUBNET MAKSK *****************
func extractSubnetMask(output string) string {
	var lines = strings.Split(output, "\n")
	
	var subNet = strings.Split(strings.TrimSpace(lines[20]), ":")[1]
	subNet = strings.TrimSpace(subNet)
	if subNet != "" {
		return subNet
	}
	return ""
}

func GetSubNetMask() string {
	return getNetInfo(
		Net_commands["SubnetMask"],
		extractSubnetMask,
	)
}
