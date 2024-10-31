package bioMetrics

import (
	"fmt"

	"github.com/AarizZafar/Nexus_protocol.git/BMmodel"
	"github.com/AarizZafar/Nexus_protocol.git/bioMetrics/controls"
)

// the below variable will contains all the biometrics values
var BMstruct BMmodel.SysBioMetrix

func ExtractSysBioMetrix() {
	BMstruct.SSID                                       = controls.GetNetssid()
	BMstruct.MAC 										= controls.GetMACAdd()
	BMstruct.SystemSerialNumber                         = controls.GetSystemSerialNumber()
	BMstruct.UUID             		        			= controls.GetSystemUUID()

	fmt.Printf("%-25s: %s\n", "SSID                      :", BMstruct.SSID)
	fmt.Printf("%-25s: %s\n", "MAC                       :", BMstruct.MAC)
	fmt.Printf("%-25s: %s\n", "System Serial Number      :", BMstruct.SystemSerialNumber)
	fmt.Printf("%-25s: %s\n", "System UUID               :", BMstruct.UUID)
}

var NBMstruct BMmodel.NetBioMetrix

func ExtractNetBioMetrix() {
	NBMstruct.SSID 										= controls.GetNetSSID()
	NBMstruct.BSSID 									= controls.GetNetBSSID()
	NBMstruct.SubNetMask 								= controls.GetSubNetMask()
	NBMstruct.NetInterfaceMAC                           = controls.GetNetInterfaceMac()

	fmt.Printf("%-25s: %s\n", "SSID                      :", NBMstruct.SSID)
	fmt.Printf("%-25s: %s\n", "BSSID                     :", NBMstruct.BSSID)
	fmt.Printf("%-25s: %s\n", "Subnet Mask               :", NBMstruct.SubNetMask)
	fmt.Printf("%-25s: %s\n", "NetInterfaceMAC           :", NBMstruct.NetInterfaceMAC)
}

