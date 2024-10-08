package bioMetrics

import (
	"fmt"

	"github.com/AarizZafar/Nexus_protocol.git/BMmodel"
	"github.com/AarizZafar/Nexus_protocol.git/bioMetrics/controls"
)

// the below variable will contains all the biometrics values
var BMstruct BMmodel.SysBioMetrix

func ExtractSysBioMetrix() {
	BMstruct.MAC 					= controls.GetMACAdd()
	BMstruct.CPUSerial 				= controls.GetCPUSerial()
	BMstruct.HardDriveSerial 		= controls.GetHardDriveSerial()
	BMstruct.MotherBoardSerial 		= controls.GetMotherBoardSerial()
	BMstruct.BIOSSerial 			= controls.GetBIOSSerial()
	BMstruct.SSDSerial 				= controls.GetSSDSerial()
	BMstruct.TPMChipID 				= controls.GetTPMChipID()
	BMstruct.GPUSerial 				= controls.GetGPUSerial()
	BMstruct.RAMSerial 				= controls.GetRAMSerial()
	BMstruct.NICID					= controls.GetNICID()
	BMstruct.BaseBoardProduct       = controls.GetBaseBoardProduct()
	BMstruct.SystemUUID             = controls.GetSystemUUID()
	BMstruct.OSInstallationID 		= controls.GetOSIntsllationID()
	BMstruct.DiskVolumeSerialNumber = controls.GetDiskVolumeSerialNo()
	BMstruct.BootROMVersion         = controls.GetBootROMVersion()
	BMstruct.GPUVendorID            = controls.GetGPUVendorID()
	BMstruct.DeviceTreeIdentifier   = controls.GetDeviceTreeIdentifier()
	BMstruct.UEFIFirmwareVersion    = controls.GetUEFIFirmwareVersion()

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
