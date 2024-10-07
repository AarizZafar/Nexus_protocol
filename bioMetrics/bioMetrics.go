package bioMetrics

import (
	"fmt"

	"github.com/AarizZafar/Nexus_protocol.git/BMmodel"
	"github.com/AarizZafar/Nexus_protocol.git/bioMetrics/controls"
)

var BMstruct BMmodel.SysBioMetrix

func ExtractSysBioMetrix() {
	BMstruct.MAC = controls.GetMACAdd()
	BMstruct.CPUSerial = controls.GetCPUSerial()
	BMstruct.HardDriveSerial = controls.GetHardDriveSerial()
	BMstruct.MotherBoardSerial = controls.GetMotherBoardSerial()
	BMstruct.BIOSSerial = controls.GetBIOSSerial()
	BMstruct.SSDSerial = controls.GetSSDSerial()
	BMstruct.TPMChipID = controls.GetTPMChipID()
	BMstruct.GPUSerial = controls.GetGPUSerial()
	BMstruct.RAMSerial = controls.GetRAMSerial()
	BMstruct.NICID = controls.GetNICID()

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
}
