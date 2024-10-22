package commands

var Net_Linux = map[string][]string{
	"SSID":                         {"nmcli", "-t", "-f", "active,ssid", "dev", "wifi", "|", "grep", "'^yes'", "|", "cut", "-d':'", "-f2"},
	"BSSID":                        {"iwconfig", "wlan0", "|", "grep", "'Access Point'"}, // replace "wlan0" with interface name
	"PublicIPAddress":              {"curl", "ifconfig.me"},
	"SubnetMask":                   {"ifconfig", "|", "grep", "-i", "'mask'"},
	"DefaultGateway":               {"ip", "route", "|", "grep", "default"},
	"NetworkInterfaceMAC":          {"ip", "link", "show"},              // or use ifconfig
	"NetworkSecurityProtocol":      {"nmcli", "device", "wifi", "list"}, // check under "SECURITY"
	"RouterConfigurationID":        {"Access router admin panel", "192.168.1.1"},
}

var Net_darwin = map[string][]string{
	"SSID":                         {"/System/Library/PrivateFrameworks/Apple80211.framework/Versions/Current/Resources/airport", "-I", "|", "grep", "' SSID'"},
	"BSSID":                        {"/System/Library/PrivateFrameworks/Apple80211.framework/Versions/Current/Resources/airport", "-I", "|", "grep", "' BSSID'"}, // Access Point MAC
	"PublicIPAddress":              {"curl", "ifconfig.me"},
	"SubnetMask":                   {"ifconfig", "|", "grep", "-i", "'netmask'"},
	"DefaultGateway":               {"netstat", "-nr", "|", "grep", "default"},
	"NetworkInterfaceMAC":          {"ifconfig"},
	"NetworkSecurityProtocol":      {"/System/Library/PrivateFrameworks/Apple80211.framework/Versions/Current/Resources/airport", "-I", "|", "grep", "'link auth'"},
}

var Net_Windows = map[string][]string{
	"SSID":                         {"netsh", "wlan", "show", "interfaces"},
	"BSSID":                        {"netsh", "wlan", "show", "interfaces"},  // Access Point MAC
	"SubnetMask":                   {"ipconfig"},
	"IPV4":                         {"ipconfig"},
	"IPV6":                         {"ipconfig"},
	"Active_NetworkInterfaceMAC":   {"getmac"},                               // or use ipconfig /all
	"Inactive_NetworkInterfaceMAC": {"getmac"},                               // or use ipconfig /all
	"NetworkSecurityProtocol":      {"netsh", "wlan", "show", "interfaces"},  // check for "Authentication"
	"RouterConfigurationID":        {"Access router admin panel", "192.168.1.1"},
}
