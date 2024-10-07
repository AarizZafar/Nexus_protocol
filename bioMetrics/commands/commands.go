package commands

var Linux = map[string][]string{
	"MAC":               {"ifconfig"},
	"CPUSerial":         {"cat", "/proc/cpuinfo"},                       // cat /proc/cpuinfo
	"HardDriveSerial":   {"lsblk", "-d", "-o", "SERIAL"},                // lsblk -d -o SERIAL
	"MotherBoardSerial": {"dmidecode", "-s", "baseboard-serial-number"}, // dmidecode -s baseboard-serial-number
	"BIOSSerial":        {"dmidecode", "-s", "system-serial-number"},    // dmidecode -s system-serial-number
	"SSDSerial":         {"lsblk", "-d", "-o", "SERIAL"},                // lsblk -d -o SERIAL
	"TPMChipID":         {"dmesg"},                                      // dmesg
	"GPUSerial":         {"lspci", "-vnn"},
	"RAMSerial":         {"sudo", "dmidecode", "-t", "memory"},
	"NICID":             {"lspci", "-nn"},
}

var Darwin = map[string][]string{
	"MAC":               {"ifconfig"},
	"CPUSerial":         {"systemctl", "-n", "machdep.cpu.serial_number"},
	"HardDriveSerial":   {"diskutil", "info", "/"},
	"MotherBoardSerial": nil,
	"BIOSSerial":        nil,
	"SSDSerial":         {"diskutil", "info", "/"},
	"TPMChipID":         nil,
	"GPUSerial":         {"system_profiler", "SPDisplaysDataType"},
	"RAMSerial":         nil,
	"NICID":             {"system_profile", "SPNetworkDataType"},
}

var Windows = map[string][]string{
	"MAC":               {"getmac"},
	"CPUSerial":         {"wmic", "cpu", "get", "ProcessorId"},                                     // wmic cpu get ProcessorId
	"HardDriveSerial":   {"wmic", "diskdrive", "get", "SerialNumber"},                              // wmic diskdrive get SerialNumber
	"MotherBoardSerial": {"wmic", "baseboard", "get", "SerialNumber"},                              // wmic baseboard get SerialNumber
	"BIOSSerial":        {"wmic", "bios", "get", "SerialNumber"},                                   // wmic bios get SerialNumber
	"SSDSerial":         {"wmic", "diskdrive", "get", "SerialNumber"},                              // wmic diskdrive get SerialNumber
	"TPMChipID":         {"wmic", "baseboard", "get", "Product,Manufacturer,Version,SerialNumber"}, // wmic baseboard get Product,Manufacturer,Version,SerialNumber
	"GPUSerial":         {"wmic", "path", "win32_videocontroller", "get", "name"},
	"RAMSerial":         {"wmic", "memorychip", "get", "SerialNumber"},
	"NICID":             {"wmic", "nic", "get", "Name"},
}
