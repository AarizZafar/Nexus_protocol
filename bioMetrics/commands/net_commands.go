package commands

var Net_Linux = map[string][]string{
	"SSID":                {"nmcli", "-t", "-f", "active,ssid", "dev", "wifi", "|", "grep", "'^yes'", "|", "cut", "-d':'", "-f2"},
	"BSSID":               {"iwconfig", "wlan0", "|", "grep", "'Access Point'"}, 
	"NetworkInterfaceMAC": {"ip", "link", "show"},              
	"SubnetMask":          {"ifconfig", "|", "grep", "-i", "'mask'"},
}

var Net_darwin = map[string][]string{
	"SSID":                {"/System/Library/PrivateFrameworks/Apple80211.framework/Versions/Current/Resources/airport", "-I", "|", "grep", "' SSID'"},
	"BSSID":               {"/System/Library/PrivateFrameworks/Apple80211.framework/Versions/Current/Resources/airport", "-I", "|", "grep", "' BSSID'"}, 
	"NetworkInterfaceMAC": {"ifconfig"},
	"SubnetMask":          {"ifconfig", "|", "grep", "-i", "'netmask'"},
}

var Net_Windows = map[string][]string{
	"SSID":                {"netsh", "wlan", "show", "interfaces"},
	"BSSID":               {"netsh", "wlan", "show", "interfaces"},  
	"NetworkInterfaceMAC": {"getmac"},                               
	"SubnetMask":          {"ipconfig"},
}