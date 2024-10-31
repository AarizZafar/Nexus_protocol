package commands

var Linux = map[string][]string{
	"SSID":                          {"nmcli", "-t", "-f", "active,ssid", "dev", "wifi", "|", "grep", "'^yes'", "|", "cut", "-d':'", "-f2"},
	"MACAddress":                    {"cat", "/sys/class/net/eth0/address"},                     // Primary MAC address
	"SystemSerialNumber":            {"sudo", "dmidecode", "-s", "system-serial-number"},        // System serial number
	"UUID":                          {"sudo", "dmidecode", "-s", "system-uuid"},                 // UUID of the system
}

var Darwin = map[string][]string{
	"SSID":                          {"/System/Library/PrivateFrameworks/Apple80211.framework/Versions/Current/Resources/airport", "-I", "|", "grep", "' SSID'"},
	"MACAddress":                    {"ifconfig", "en0", "|", "grep", "ether"},                        // MAC address of the primary interface
	"SystemSerialNumber":            {"system_profiler", "SPHardwareDataType", "|", "grep", "Serial"}, // System serial number
	"UUID":                          {"ioreg", "-l", "|", "grep", "IOPlatformUUID"},                   // UUID of the system
}

var Windows = map[string][]string{
	"SSID":                          {"netsh", "wlan", "show", "interfaces"},
	"MACAddress":                    {"getmac"},
	"SystemSerialNumber":            {"wmic", "bios", "get", "serialnumber"},                                  // wmic bios get serialnumber
	"UUID":                          {"wmic", "csproduct", "get", "uuid"},                                     // wmic csproduct get uuid
}
