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
	BMstruct.CPUSerial 									= controls.GetCPUSerial()
	BMstruct.HardDriveSerial 							= controls.GetHardDriveSerial()
	BMstruct.MotherBoardSerial 							= controls.GetMotherBoardSerial()
	BMstruct.BIOSSerial 								= controls.GetBIOSSerial()
	BMstruct.SSDSerial 									= controls.GetSSDSerial()
	BMstruct.TPMChipID 									= controls.GetTPMChipID()
	BMstruct.GPUSerial 									= controls.GetGPUSerial()
	BMstruct.RAMSerial 									= controls.GetRAMSerial()
	BMstruct.NICID										= controls.GetNICID()
	BMstruct.BaseBoardProduct       					= controls.GetBaseBoardProduct()
	BMstruct.SystemUUID             					= controls.GetSystemUUID()
	BMstruct.OSInstallationID 							= controls.GetOSIntsllationID()
	BMstruct.DiskVolumeSerialNumber 					= controls.GetDiskVolumeSerialNo()
	BMstruct.BootROMVersion         					= controls.GetBootROMVersion()
	BMstruct.GPUVendorID            					= controls.GetGPUVendorID()
	BMstruct.DeviceTreeIdentifier   					= controls.GetDeviceTreeIdentifier()
	BMstruct.UEFIFirmwareVersion    					= controls.GetUEFIFirmwareVersion()

	fmt.Printf("%-25s: %s\n", "SSID                      :", BMstruct.SSID)
	fmt.Printf("%-25s: %s\n", "MAC                       :", BMstruct.MAC)
	fmt.Printf("%-25s: %s\n", "CPUSerial                 :", BMstruct.CPUSerial)
	fmt.Printf("%-25s: %s\n", "Hard Drive Serial         :", BMstruct.HardDriveSerial)
	fmt.Printf("%-25s: %s\n", "Mother Board Serial       :", BMstruct.MotherBoardSerial)
	fmt.Printf("%-25s: %s\n", "BIOS Serial               :", BMstruct.BIOSSerial)
	fmt.Printf("%-25s: %s\n", "SSD Serial                :", BMstruct.SSDSerial)
	fmt.Printf("%-25s: %s\n", "TPM chip ID               :", BMstruct.TPMChipID)
	fmt.Printf("%-25s: %s\n", "GPU Serial                :", BMstruct.GPUSerial)
	fmt.Printf("%-25s: %s\n", "RAM Serial                :", BMstruct.RAMSerial)
	fmt.Printf("%-25s: %s\n", "NICID                     :", BMstruct.NICID)
	fmt.Printf("%-25s: %s\n", "Base Board Product        :", BMstruct.BaseBoardProduct)
	fmt.Printf("%-25s: %s\n", "System UUID               :", BMstruct.SystemUUID)
	fmt.Printf("%-25s: %s\n", "OS Installation ID        :", BMstruct.OSInstallationID)
	fmt.Printf("%-25s: %s\n", "Disk Volume Serial no     :", BMstruct.DiskVolumeSerialNumber)
	fmt.Printf("%-25s: %s\n", "Boot rom version          :", BMstruct.BootROMVersion)
	fmt.Printf("%-25s: %s\n", "GPU Vendor ID             :", BMstruct.GPUVendorID)
	fmt.Printf("%-25s: %s\n", "Device Tree Identifier    :", BMstruct.DeviceTreeIdentifier)
	fmt.Printf("%-25s: %s\n", "UEFI Firmware Version     :", BMstruct.UEFIFirmwareVersion)
}

var NBMstruct BMmodel.NetBioMetrix

func ExtractNetBioMetrix() {
	NBMstruct.SSID 										= controls.GetNetSSID()
	NBMstruct.BSSID 									= controls.GetNetBSSID()
	NBMstruct.SubNetMask 								= controls.GetSubNetMask()
	NBMstruct.IPV4_DG        							= controls.GetIPV4_DG()
	NBMstruct.IPV6_DG        							= controls.GetIPV6_DG()
	NBMstruct.Active_MAC   								= controls.GetActiveMAC()
	NBMstruct.Inactive_MAC   							= controls.GetInActiveMAC()
	NBMstruct.Security_proto 							= controls.GetNetSecProtocol()

	fmt.Printf("%-25s: %s\n", "SSID                      :", NBMstruct.SSID)
	fmt.Printf("%-25s: %s\n", "BSSID                     :", NBMstruct.BSSID)
	fmt.Printf("%-25s: %s\n", "Subnet Mask               :", NBMstruct.SubNetMask)
	fmt.Printf("%-25s: %s\n", "IPV4 (Default Gateway)    :", NBMstruct.IPV4_DG)
	fmt.Printf("%-25s: %s\n", "IPV6 (Default Gateway)    :", NBMstruct.IPV6_DG)
	fmt.Printf("%-25s: %s\n", "Active MAC                :", NBMstruct.Active_MAC)
	fmt.Printf("%-25s: %s\n", "InActive MAC              :", NBMstruct.Inactive_MAC)
	fmt.Printf("%-25s: %s\n", "Network Security Protocol :", NBMstruct.Security_proto)
}

