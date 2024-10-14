package BMmodel

type SysBioMetrix struct {
	MAC                    string
	CPUSerial              string
	HardDriveSerial        string
	MotherBoardSerial      string
	BIOSSerial             string
	SSDSerial              string
	TPMChipID              string
	RAMSerial              string
	GPUSerial              string
	NICID                  string
	BaseBoardProduct       string
	SystemUUID             string
	OSInstallationID       string
	DiskVolumeSerialNumber string
	BootROMVersion         string
	GPUVendorID            string
	DeviceTreeIdentifier   string
	UEFIFirmwareVersion    string
}

type NetBioMetrix struct {
	SSID                   string
	BSSID                  string 
	PublicIPAdd            string 
	SubNetMask 			   string
	IPV4_DG                string
	IPV6_DG                string
	Active_MAC             string
	Inactive_MAC           string
	Security_proto         string
}
