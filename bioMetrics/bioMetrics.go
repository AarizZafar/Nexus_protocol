package bioMetrics

import (
	"fmt"
	"os/exec" // used to run external commands or progs from withing a Go program and capture their output
	"runtime" // provides information obout the go runtime system
	"strings"

	"github.com/AarizZafar/Nexus_protocol.git/BMmodel"
)

var currOS = runtime.GOOS

// Accessing it from the BMmodel file
var BMstruct BMmodel.SysBioMetrix
func AssignBMStruct() {
	BMstruct = BMmodel.BMstruct
}

// ... - (variadic parameter) accept a parameter type defined or nil, the arguments are collected in a slice of string
func safeCommandOutput(commands string, args ...string) string {
	cmd := exec.Command(commands, args...) // creates a command that will run in the OS (like running in the terminal)
	// returns a *exec.Cmd (pointer) to exec.Cmd struct -> represents the command to be run in the OS
	output, err := cmd.Output() // returns -> []byte - result of the command, error
	if err != nil {
		return ""
	}
	// strings.TrimSpace(string(output))        raw output of the command executed in terminal
	return strings.TrimSpace(string(output)) // converts the output -> string, removes any leading or trailing spaces,
}

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

func getMACAdd() string {
	var output string

	switch currOS {
	case "linux", "darwin":
		output = safeCommandOutput("ifconfig")
	case "windows":
		output = safeCommandOutput("getmac")
	}

	lines := strings.Split(output, "\n") // splitting the raw output string into a slice of string based on spaces
	parts := strings.Fields(lines[2])    // return type is a slice of string
	if len(parts) > 0 {
		return parts[0]
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

func getCPUSerial() string {
	return getSysInfo(
		// cat /proc/cpuinfo
		[]string{"cat", "/proc/cpuinfo"},
		[]string{"systemctl", "-n", "machdep.cpu.serial_number"},
		// wmic cpu get ProcessorId
		[]string{"wmic", "cpu", "get", "ProcessorId"},
		formatStr,
	)
}

func getHardDrive() string {
	return getSysInfo(
		// lsblk -d -o SERIAL
		[]string{"lsblk", "-d", "-o", "SERIAL"},
		[]string{"diskutil", "info", "/"},
		// wmic diskdrive get SerialNumber
		[]string{"wmic", "diskdrive", "get", "SerialNumber"},
		formatStr,
	)
}

func getMotherBoardSerial() string {
	return getSysInfo(
		// dmidecode -s baseboard-serial-number
		[]string{"dmidecode", "-s", "baseboard-serial-number"},
		nil,
		// wmic baseboard get SerialNumber
		[]string{"wmic", "baseboard", "get", "SerialNumber"},
		formatStr,
	)
}

func getBIOSSerial() string {
	return getSysInfo(
		// dmidecode -s system-serial-number
		[]string{"dmidecode", "-s", "system-serial-number"},
		nil,
		// wmic bios get SerialNumber
		[]string{"wmic", "bios", "get", "SerialNumber"},
		formatStr,
	)
}

func getSSDSerial() string {
	return getSysInfo(
		// lsblk -d -o SERIAL
		[]string{"lsblk","-d","-o","SERIAL"},
		[]string{"diskutil","info","/"},
		// wmic diskdrive get SerialNumber
		[]string{"wmic","diskdrive","get","SerialNumber"},
		formatStr,
	)
}

// ***************************************************************************

func extractTPMSerial(output string) string {
	lines := strings.Split(output, "\n")

	var count = 0
	var trimmed_data []string
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if count == 1 {
			data := strings.Fields(trimmed)
			trimmed_data = append(trimmed_data, data...)
		}
		count++
	}
	if trimmed_data[4] != "" {
		return trimmed_data[4]
	}
	return ""
}

func getTPMChipID() string {
	return getSysInfo(
		// dmesg
		[]string{"dmesg"}, 
		nil,               
		// wmic baseboard get Product,Manufacturer,Version,SerialNumber
		[]string{"wmic", "baseboard", "get", "Product,Manufacturer,Version,SerialNumber"}, 
		extractTPMSerial, 
	)
}

// ***************************************************************************
func extractGPUSerial(output string) string {
	lines := strings.Split(output, "\n")   // splitting based on new line 

	var count = 0
	var trimmed_data []string
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)  // trimming any extra spaces before or after the line 

		if count == 1 {
			data := strings.Fields(trimmed)  // splitting based on spaces between words
			trimmed_data = append(trimmed_data, data...)
		}
		count++
	}

	if trimmed_data[0] != "" {
		return trimmed_data[0]
	}
	return ""
}

func getGPUSerial() string {
	return getSysInfo(
		[]string{"lspci", "-vnn"},
		[]string{"system_profiler","SPDisplaysDataType"},
		[]string{"wmic", "path", "win32_videocontroller", "get", "name"},
		extractGPUSerial,
	)
}

// ***************************************************************************

func extractRAMSerial(output string) string {
	lines := strings.Split(output, "\n")

	var trimmedLines []string
	for _, line := range lines {
		trimmedLines = append(trimmedLines, strings.TrimSpace(line))
	}
	if trimmedLines[2] != "" {
		return trimmedLines[2]
	}
	return ""
}

func getRAMSerial() string {
	return getSysInfo(
		[]string{"sudo", "dmidecode", "-t", "memory"},
		nil,
		[]string{"wmic", "memorychip", "get", "SerialNumber"},
		extractRAMSerial,
	)
}

// ***************************************************************************

// ***************************************************************************

func extractNICID(output string) string {
	lines := strings.Split(output, "\n")

	var trimmed_data []string 
	var count = 0
	for _, line := range lines {
		trimmed_lines := strings.TrimSpace(line)

		if count == 2 {
			data := strings.Fields(trimmed_lines)
			trimmed_data = append(trimmed_data, data...)
		}
		count++;
	}
	if trimmed_data[3] != "" {
		return trimmed_data[3]
	}

	return ""
}

func getNICID() string {
	return getSysInfo(
		[]string{"lspci", "-nn"},
		[]string{"system_profile","SPNetworkDataType"},
		[]string{"wmic","nic","get","Name"},
		extractNICID,
	)
}

// ***************************************************************************



func ExtSysBioMetrix() {
	BMstruct.MAC = getMACAdd()
	BMstruct.CPUSerial = getCPUSerial()
	BMstruct.HardDriveSerial = getHardDrive()
	BMstruct.MotherBoardSerial = getMotherBoardSerial()
	BMstruct.BIOSSerial = getBIOSSerial()
	BMstruct.SSDSerial = getSSDSerial()
	BMstruct.TPMChipID = getTPMChipID()
	BMstruct.RAMSerial = getRAMSerial()
	BMstruct.GPUSerial = getGPUSerial()
	BMstruct.NICID = getNICID()

	fmt.Println("MAC ADD : ", BMstruct.MAC)
	fmt.Println("CPUSerial : ", BMstruct.CPUSerial)
	fmt.Println("HardDrive Serial : ", BMstruct.HardDriveSerial)
	fmt.Println("Mother Board Serial : ", BMstruct.MotherBoardSerial)
	fmt.Println("BIOS Serial number : ", BMstruct.BIOSSerial)
	fmt.Println("SSD Serial number : ", BMstruct.SSDSerial)
	fmt.Println("TPM Chip ID : ", BMstruct.TPMChipID)
	fmt.Println("GPU Serial number : ", BMstruct.GPUSerial)
	fmt.Println("RAM Serial number : ", BMstruct.RAMSerial)
	fmt.Println("NICID : ", BMstruct.NICID)
}
