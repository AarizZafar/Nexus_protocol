package controls

import (
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

func getSysInfo(windowsCmd []string, formatStr func(string) string) string {
	var output string
	output = safeCommandOutput(windowsCmd[0], windowsCmd[1:]...)
	if formatStr != nil {
		output = formatStr(output)
	}
	return output
}

// ********************* MAC ********************
func GetMACAdd() string {
	var output = safeCommandOutput(executeCmd["MAC"][0], executeCmd["MAC"][1:]...)

	lines := strings.Split(output, "\n") // splitting the raw output string into a slice of string based on spaces
	parts := strings.Fields(lines[2])    // return type is a slice of string
	if len(parts) > 0 {
		return parts[0]
	}
	return ""
}

// ***************** CPU SERIAL *****************
func GetCPUSerial() string {
	return getSysInfo(
		executeCmd["CPUSerial"],
		formatStr,
	)
}

// *********** HARD DRIVE SERIAL ****************
func GetHardDriveSerial() string {
	return getSysInfo(
		executeCmd["HardDriveSerial"],
		formatStr,
	)
}

// ********** MOTHER BOARD SERIAL **************
func GetMotherBoardSerial() string {
	return getSysInfo(
		executeCmd["MotherBoardSerial"],
		formatStr,
	)
}

// *********** BIOS Serial *********************
func GetBIOSSerial() string {
	return getSysInfo(
		executeCmd["BIOSSerial"],
		formatStr,
	)
}

// *********** SSD Serial **********************
func GetSSDSerial() string {
	return getSysInfo(
		executeCmd["SSDSerial"],
		formatStr,
	)
}

// *********** TPM Serial **********************
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

func GetTPMChipID() string {
	return getSysInfo(
		executeCmd["TPMChipID"],
		extractTPMSerial, 
	)
}

// *********** GPU Serial **********************
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

func GetGPUSerial() string {
	return getSysInfo(
		executeCmd["GPUSerial"],
		extractGPUSerial,
	)
}

// *********** RAM Serial **********************
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

func GetRAMSerial() string {
	return getSysInfo(
		executeCmd["RAMSerial"],
		extractRAMSerial,
	)
}

// **************** NICID **********************
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

func GetNICID() string {
	return getSysInfo(
		executeCmd["NICID"],
		extractNICID,
	)
}

// ******** BASE BOARD PRODUCT ID **************
func extractBoardProduct(output string) string {
	lines := strings.Split(output, "\n")
	// var trimmed_data []string 
	
	var count = 0
	var trimmed_data []string
	for _, line := range lines {
		trimmed_line := strings.TrimSpace(line)    
		
		if count == 1 {
			data := strings.Fields(trimmed_line)
			trimmed_data = append(trimmed_data, data...)
		}
		count++;
	}
	if trimmed_data[0] != "" {
		return trimmed_data[0]
	}
	
	return ""
}

func GetBaseBoardProduct() string {
	return getSysInfo(
		executeCmd["BaseboardProductID"],
		extractBoardProduct,
	)
}

// ************* SYSTEM UUID *******************
func extractSystemUUID(output string) string {
	lines := strings.Split(output, "\n")
	
	var count = 0
	var trimmed_data []string
	for _, line := range lines {
		trimmed_line := strings.TrimSpace(line)     
		
		if count == 1 {
			data := strings.Fields(trimmed_line)
			trimmed_data = append(trimmed_data, data...)
		}
		count++;
	}
	if trimmed_data[0] != "" {
		return trimmed_data[0]
	}
	
	return ""
}

func GetSystemUUID() string {
	return getSysInfo(
		executeCmd["SystemUUID"],
		extractSystemUUID,
	)
}

// ********** OS INSTALLATION ID ***************
func extractOSID(output string) string {
	lines := strings.Split(output, "\n")
	
	var count = 0
	var trimmed_data []string
	for _, line := range lines {
		trimmed_line := strings.TrimSpace(line)     
		
		if count == 1 {
			data := strings.Fields(trimmed_line)
			trimmed_data = append(trimmed_data, data...)
		}
		count++;
	}
	if trimmed_data[0] != "" {
		return trimmed_data[0]
	}
	
	return ""
}

func GetOSIntsllationID() string {
	return getSysInfo(
		executeCmd["OSInstallationID"],
		extractOSID,
	)
}

// ****** DISK VOLUME SERIAL NUMBER ************
func extractDiskVolumeSerilNo(output string) string {
	lines := strings.Split(output, "\n")
	
	var count = 0
	var trimmed_data []string
	for _, line := range lines {
		trimmed_line := strings.TrimSpace(line)     
		
		if count >= 1 {
			data := strings.Fields(trimmed_line)
			trimmed_data = append(trimmed_data, data...)
		}
		count++;
	}
	if trimmed_data[0] != "" {
		return trimmed_data[0] + trimmed_data[1] + trimmed_data[2]
	}
	return ""
}

func GetDiskVolumeSerialNo() string {
	return getSysInfo(
		executeCmd["DiskVolumeSerialNumber"],
		extractDiskVolumeSerilNo,
	)
} 

// ************* BOOT ROM VERSION **************
func extractBootROMVersion(output string) string {
	lines := strings.Split(output, "\n")
	
	var count = 0
	var trimmed_data []string
	for _, line := range lines {
		trimmed_line := strings.TrimSpace(line)     
		
		if count == 1 {
			data := strings.Fields(trimmed_line)
			trimmed_data = append(trimmed_data, data...)
		}
		count++;
	}
	if trimmed_data[2] != "" {
		return trimmed_data[2]
	}
	return ""
}

func GetBootROMVersion() string {
	return getSysInfo(
		executeCmd["BootROMVersion"],
		extractBootROMVersion,
	)
}

// ************* GPU VENDOR ID ******************
func extractGPUVendorID(output string) string {
	lines := strings.Split(output, "\n")
	
	var count = 0
	var trimmed_data []string
	for _, line := range lines {
		trimmed_line := strings.TrimSpace(line)    
		
		if count == 1 {
			data := strings.Fields(trimmed_line)
			trimmed_data = append(trimmed_data, data...)
		}
		count++;
	}
	if trimmed_data[0] != "" {
		return trimmed_data[0]
	}
	return ""
}

func GetGPUVendorID() string {
	return getSysInfo(
		executeCmd["GPUVendorID"],
		extractGPUVendorID,
	)
}

// ******** DEVICE TREE IDENTIFIER **************
func extractDeviceTreeIdentifier(output string) string {
	lines := strings.Split(output, "\n")
	
	var count = 0
	var trimmed_data []string
	for _, line := range lines {
		trimmed_line := strings.TrimSpace(line)     
		
		if count == 1 {
			data := strings.Fields(trimmed_line)
			trimmed_data = append(trimmed_data, data...)
		}
		count++;
	}
	if trimmed_data[1] != "" {
		return trimmed_data[1]
	}
	
	return ""
}

func GetDeviceTreeIdentifier() string {
	return getSysInfo(
		executeCmd["DeviceTreeIdentifier"],
		extractDeviceTreeIdentifier,
	)
}

// ******** UEFI FIRMWARE VERSION ***************
func extractUEFIFirmwareVersion(output string) string {
	lines := strings.Split(output, "\n")

	var count = 0
	var trimmed_data []string
	for _, line := range lines {
		trimmed_line := strings.TrimSpace(line)     

		if count == 1 {
			data := strings.Fields(trimmed_line)
			trimmed_data = append(trimmed_data, data...)
		}
		count++;
	}
	if trimmed_data[0] != "" {
		return trimmed_data[0]
	}

	return ""
}

func GetUEFIFirmwareVersion() string {
	return getSysInfo(
		executeCmd["UEFIFirmwareVersion"],
		extractUEFIFirmwareVersion,
	)
}










