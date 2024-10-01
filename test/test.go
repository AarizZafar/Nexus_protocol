package main

import (
	"os/exec"
	"runtime"
	"strings"
)

var currOS = runtime.GOOS

// Helper function to execute system command and return the cleaned output
func safeCommandOutput(command string, args ...string) string {
	cmd := exec.Command(command, args...)
	output, err := cmd.Output()
	if err != nil {
		return ""
	}
	return strings.TrimSpace(string(output))
}

// Function to extract system information and clean unnecessary labels
func formatStr(output string) string {
	lines := strings.Split(output, "\n")
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if trimmed != "" && !strings.Contains(trimmed, "ProcessorId") && !strings.Contains(trimmed, "SerialNumber") && !strings.Contains(trimmed, "Name") {
			return trimmed
		}
	}
	return ""
}

func getSysInfo(linuxCmd []string, darwinCmd []string, windowsCmd []string, formatStr func(string) string) string {
	var output string
	switch currOS {
	case "linux":
		output = safeCommandOutput(linuxCmd[0], linuxCmd[1:]...)
	case "darwin":
		output = safeCommandOutput(darwinCmd[0], darwinCmd[1:]...)
	case "windows":
		output = safeCommandOutput(windowsCmd[0], windowsCmd[1:]...)
	default:
		return ""
	}
	if formatStr != nil {
		output = formatStr(output)
	}
	return output
}

// Function to get BIOS Serial Number
func getBIOSSerial() string {
	return getSysInfo(
		[]string{"dmidecode", "-s", "system-serial-number"},
		nil,
		[]string{"wmic", "bios", "get", "SerialNumber"},
		formatStr,
	)
}

// Function to get SSD Serial Number
func getSSDSerial() string {
	return getSysInfo(
		[]string{"lsblk", "-d", "-o", "SERIAL"},
		[]string{"diskutil", "info", "/"},
		[]string{"wmic", "diskdrive", "get", "SerialNumber"},
		formatStr,
	)
}

// Function to get TPM ID
func getTPMChipID() string {
	return getSysInfo(
		[]string{"dmesg"},
		nil,
		[]string{"wmic", "tpminformation", "get", "ManufacturerID"},
		formatStr,
	)
}

// Function to get GPU Serial Number
func getGPUSerial() string {
	return getSysInfo(
		[]string{"lspci", "-vnn"},
		[]string{"system_profiler", "SPDisplaysDataType"},
		[]string{"wmic", "path", "win32_videocontroller", "get", "name"},
		formatStr,
	)
}

// Function to get RAM Serial Number
func getRAMSerial() string {
	return getSysInfo(
		[]string{"sudo", "dmidecode", "-t", "memory"},
		nil,
		[]string{"wmic", "memorychip", "get", "SerialNumber"},
		formatStr,
	)
}

// Function to get NIC ID
func getNICID() string {
	return getSysInfo(
		[]string{"lspci", "-nn"},
		[]string{"system_profiler", "SPNetworkDataType"},
		[]string{"wmic", "nic", "get", "Name"},
		formatStr,
	)
}

// Example function to call and display all biometrics
func AssignBiometrics() {
	biosSerial := getBIOSSerial()
	ssdSerial := getSSDSerial()
	tpmID := getTPMChipID()
	gpuSerial := getGPUSerial()
	ramSerial := getRAMSerial()
	nicID := getNICID()

	// Print the biometrics
	println("BIOS Serial Number:", biosSerial)
	println("SSD Serial Number:", ssdSerial)
	println("TPM ID:", tpmID)
	println("GPU Serial Number:", gpuSerial)
	println("RAM Serial Number:", ramSerial)
	println("NIC ID:", nicID)
}

func main() {
	AssignBiometrics()
}