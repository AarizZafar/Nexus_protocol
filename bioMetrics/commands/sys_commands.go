package commands

var Linux = map[string][]string{
	"SSID":                   {"nmcli", "-t", "-f", "active,ssid", "dev", "wifi", "|", "grep", "'^yes'", "|", "cut", "-d':'", "-f2"},
	"MAC":                    {"ifconfig"},
	"CPUSerial":              {"cat", "/proc/cpuinfo"},                       // cat /proc/cpuinfo
	"HardDriveSerial":        {"lsblk", "-d", "-o", "SERIAL"},                // lsblk -d -o SERIAL
	"MotherBoardSerial":      {"dmidecode", "-s", "baseboard-serial-number"}, // dmidecode -s baseboard-serial-number
	"BIOSSerial":             {"dmidecode", "-s", "system-serial-number"},    // dmidecode -s system-serial-number
	"SSDSerial":              {"lsblk", "-d", "-o", "SERIAL"},                // lsblk -d -o SERIAL
	"TPMChipID":              {"dmesg"},                                      // dmesg
	"GPUSerial":              {"lspci", "-vnn"},
	"RAMSerial":              {"sudo", "dmidecode", "-t", "memory"},
	"NICID":                  {"lspci", "-nn"},
	"BaseboardProductID":     {"dmidecode", "-s", "baseboard-product-name"},
	"SystemUUID":             {"dmidecode", "-s", "system-uuid"},
	"OSInstallationID":       {"cat", "/var/lib/dbus/machine-id"},
	"DiskVolumeSerialNumber": {"lsblk", "-o", "UUID"},
	"BootROMVersion":         {"dmidecode", "-s", "bios-version"},
	"GPUVendorID":            {"lspci", "-nn", "|", "grep", "VGA"},
	"DeviceTreeIdentifier":   {"cat", "/proc/device-tree/model"}, // (Linux ARM-based only)
	"UEFIFirmwareVersion":    {"cat", "/sys/firmware/efi/fw_platform_size"},
}

var Darwin = map[string][]string{
	"SSID":                    {"/System/Library/PrivateFrameworks/Apple80211.framework/Versions/Current/Resources/airport", "-I", "|", "grep", "' SSID'"},
	"MAC":                    {"ifconfig"},
	"CPUSerial":              {"systemctl", "-n", "machdep.cpu.serial_number"},
	"HardDriveSerial":        {"diskutil", "info", "/"},
	"MotherBoardSerial":      nil,
	"BIOSSerial":             nil,
	"SSDSerial":              {"diskutil", "info", "/"},
	"TPMChipID":              nil,
	"GPUSerial":              {"system_profiler", "SPDisplaysDataType"},
	"RAMSerial":              nil,
	"NICID":                  {"system_profile", "SPNetworkDataType"},
	"BaseboardProductID":     {"system_profiler", "SPHardwareDataType"},
	"SystemUUID":             {"ioreg", "-l", "|", "grep", "IOPlatformUUID"},
	"OSInstallationID":       {"ioreg", "-l", "|", "grep", "IOPlatformUUID"},
	"DiskVolumeSerialNumber": {"diskutil", "info", "/"},
	"BootROMVersion":         {"system_profiler", "SPHardwareDataType"},
	"GPUVendorID":            {"system_profiler", "SPDisplaysDataType"},
	"DeviceTreeIdentifier":   {"ioreg", "-l", "|", "grep", "device-tree"},
	"UEFIFirmwareVersion":    {"system_profiler", "SPHardwareDataType"},
}

var Windows = map[string][]string{
	"SSID":                   {"netsh", "wlan", "show", "interfaces"},
	"MAC":                    {"getmac"},
	"CPUSerial":              {"wmic", "cpu", "get", "ProcessorId"},                                     // wmic cpu get ProcessorId
	"HardDriveSerial":        {"wmic", "diskdrive", "get", "SerialNumber"},                              // wmic diskdrive get SerialNumber
	"MotherBoardSerial":      {"wmic", "baseboard", "get", "SerialNumber"},                              // wmic baseboard get SerialNumber
	"BIOSSerial":             {"wmic", "bios", "get", "SerialNumber"},                                   // wmic bios get SerialNumber
	"SSDSerial":              {"wmic", "diskdrive", "get", "SerialNumber"},                              // wmic diskdrive get SerialNumber
	"TPMChipID":              {"wmic", "baseboard", "get", "Product,Manufacturer,Version,SerialNumber"}, // wmic baseboard get Product,Manufacturer,Version,SerialNumber
	"GPUSerial":              {"wmic", "path", "win32_videocontroller", "get", "name"},
	"RAMSerial":              {"wmic", "memorychip", "get", "SerialNumber"},
	"NICID":                  {"wmic", "nic", "get", "Name"},
	"BaseboardProductID":     {"wmic", "baseboard", "get", "Product"},
	"SystemUUID":             {"wmic", "csproduct", "get", "UUID"},
	"OSInstallationID":       {"wmic", "os", "get", "SerialNumber"},
	"DiskVolumeSerialNumber": {"wmic", "volume", "get", "SerialNumber"},
	"BootROMVersion":         {"wmic", "bios", "get", "Version"},
	"GPUVendorID":            {"wmic", "path", "win32_videocontroller", "get", "PNPDeviceID"},
	"DeviceTreeIdentifier":   {"wmic", "computersystem", "get", "Model"}, // wmic computersystem get Model
	"UEFIFirmwareVersion":    {"powershell", "(Get-WmiObject -Class Win32_BIOS).SMBIOSBIOSVersion"},
}
