package controls

import (
	"fmt"
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
	// fmt.Println(command , "", args)
	cmd := exec.Command(command, args...)   // cannot take sheel features like the |(pipe)
	output, err := cmd.Output()
	if err != nil {
		return ""
	}
	return strings.TrimSpace(string(output))
}

func getNetInfo(command []string, formatStr func(string) string) string {
	var output string
	// fmt.Println(command[1:])
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

// **************** PUBLIC IP ADD ****************
func extractPublicIP(output string) string {
	var lines = strings.Split(output, "\n")
	var PublicIP = lines[0]
	
	if PublicIP != "" {
		return PublicIP
	}
	return ""
}

func GetNetPublicIP() string {
	return getNetInfo(
		Net_commands["PublicIPAddress"],
		extractPublicIP,
		) 
	}
	
// ***************** SUBNET MAKS *****************
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

// ******************** IPV4 *********************
func extractIPV4(output string) string {
	var lines = strings.Split(output, "\n")
	
	var IPV4 = strings.SplitN(strings.TrimSpace(lines[21]), ":", 2)[1]
	if IPV4 != "" {
		return strings.TrimSpace(IPV4)
	}
	return ""
}

func GetIPV4_DG() string {
	fmt.Println("------") //----------------------------------------------------------------------------------------------------
	return getNetInfo(
		Net_commands["IPV4"],
		extractIPV4,
	)
}

// ******************** IPV6 *********************
func extractIPV6(output string) string{
	var lines = strings.Split(output, "\n")
	
	var IPV6 = strings.TrimSpace(lines[22])
	if IPV6 != "" {
		return IPV6
	}
	return  ""
}

func GetIPV6_DG() string {
	return getNetInfo(
		Net_commands["IPV6"],
		extractIPV6,
	)
}

// ***************** ACTIVE MAC ******************
func extractActiveMac(output string) string{
	var lines = strings.Split(output, "\n")
	
	var activeMac = strings.Split(strings.TrimSpace(lines[2]), " ")
	if activeMac[0] != "" {
		return activeMac[0]
	}
	return ""
}

func GetActiveMAC() string {
	return getNetInfo(
		Net_commands["Active_NetworkInterfaceMAC"],
		extractActiveMac,
	)
}

// *************** INACTIVE MAC ******************
func extractInActiveMac(output string) string {
	var lines = strings.Split(output, "\n")
	
	var inactiveMac = strings.Split(strings.TrimSpace(lines[3]), " ")
	if inactiveMac[0] != "" {
		return inactiveMac[0]
	}
	return ""
}

func GetInActiveMAC() string {
	return getNetInfo(
		Net_commands["Inactive_NetworkInterfaceMAC"],
		extractInActiveMac,
	)
}

// ******** NETWORK SECURITY PROTOCOL ************
func extractNetSecProto(output string) string {
	var lines = strings.Split(output, "\n")

	var SecProtocol = strings.Split(strings.TrimSpace(lines[12]),":")[1]
	
	if SecProtocol != "" {
		return strings.TrimSpace(SecProtocol)
	}
	return ""
}

func GetNetSecProtocol() string {
	return getNetInfo(
		Net_commands["NetworkSecurityProtocol"],
		extractNetSecProto,
	)
}

