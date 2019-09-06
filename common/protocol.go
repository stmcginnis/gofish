//
// SPDX-License-Identifier: BSD-3-Clause
//

package common

// Protocol is the communication protocol.
type Protocol string

const (
	// PCIeProtocol shall mean that this device conforms to the PCI-SIG PCI
	// Express Base Specification.
	PCIeProtocol Protocol = "PCIe"
	// AHCIProtocol shall mean that this device conforms to the Intel
	// Advanced Host Controller Interface Specification.
	AHCIProtocol Protocol = "AHCI"
	// UHCIProtocol shall mean that this device conforms to the Intel
	// Universal Host Controller Interface Specification, Enhanced Host
	// Controller Interface Specification, or the Extensible Host Controller
	// Interface specification.
	UHCIProtocol Protocol = "UHCI"
	// SASProtocol shall mean that this device conforms to the T10 SAS
	// Protocol Layer Specification.
	SASProtocol Protocol = "SAS"
	// SATAProtocol shall mean that this device conforms to the Serial ATA
	// International Organization Serial ATA Specification.
	SATAProtocol Protocol = "SATA"
	// USBProtocol shall mean that this device conforms to the USB
	// Implementers Forum Universal Serial Bus Specification.
	USBProtocol Protocol = "USB"
	// NVMeProtocol shall mean that this device conforms to the Non-Volatile
	// Memory Host Controller Interface Specification Specification.
	NVMeProtocol Protocol = "NVMe"
	// FCProtocol shall mean that this device conforms to the T11 Fibre
	// Channel Physical and Signaling Interface Specification.
	FCProtocol Protocol = "FC"
	// ISCSIProtocol shall mean that this device conforms to the IETF
	// Internet Small Computer Systems Interface (iSCSI) Specification.
	ISCSIProtocol Protocol = "iSCSI"
	// FCoEProtocol shall mean that this device conforms to the T11 FC-BB-5
	// Specification.
	FCoEProtocol Protocol = "FCoE"
	// FCPProtocol shall indicate the INCITS 481: Information technology -
	// Fibre Channel Protocol for SCSI.  The Fibre Channel SCSI Protocol.
	FCPProtocol Protocol = "FCP"
	// FICONProtocol shall indicate the (ANSI FC-SB-3 Single-Byte Command
	// Code Sets-3 Mapping Protocol for the Fibre Channel(FC) protocol.
	// FICON (FIbre CONnection) is the IBM proprietary name for this
	// protocol.
	FICONProtocol Protocol = "FICON"
	// NVMeOverFabricsProtocol shall mean that this device conforms to the
	// NVM Express over Fabrics Specification.
	NVMeOverFabricsProtocol Protocol = "NVMeOverFabrics"
	// SMBProtocol shall mean that this device conforms to the Microsoft
	// Server Message Block Protocol.
	SMBProtocol Protocol = "SMB"
	// NFSv3Protocol shall mean that this device conforms to the Network File
	// System protocol as defined by RFC 1813.
	NFSv3Protocol Protocol = "NFSv3"
	// NFSv4Protocol shall mean that this device conforms to the Network File
	// System protocol as defined by RFC 3010 or RFC 5661.
	NFSv4Protocol Protocol = "NFSv4"
	// HTTPProtocol shall mean that this device conforms to the Hypertext
	// Transfer protocol as defined by RFC 2068 or RFC 2616.
	HTTPProtocol Protocol = "HTTP"
	// HTTPSProtocol shall mean that this device conforms to the Hypertext
	// Transfer protocol as defined by RFC 2068 or RFC 2616 utilizing
	// Transport Layer Security as specified by RFC 5246 or RFC 6176.
	HTTPSProtocol Protocol = "HTTPS"
	// FTPProtocol shall mean that this device conforms to the File Transfer
	// protocol as defined by RFC 114.
	FTPProtocol Protocol = "FTP"
	// SFTPProtocol shall mean that this device conforms to the File Transfer
	// protocol as defined by RFC 114 utilizing Transport Layer Security as
	// specified by RFC 5246 or RFC 6176.
	SFTPProtocol Protocol = "SFTP"
	// IWARPProtocol shall mean that this device conforms to the iWARP
	// protocol as defined by RFC 5042 utilizing Transport Layer mechanisms
	// as specified by RFC 5043 or RFC 5044.
	IWARPProtocol Protocol = "iWARP"
	// RoCEProtocol shall mean that this device conforms to the RDMA over
	// Converged Ethernet protocol as defined by the Infiniband Architecture
	// Specification.
	RoCEProtocol Protocol = "RoCE"
	// RoCEv2Protocol shall mean that this device conforms to the RDMA over
	// Converged Ethernet version 2 protocol as defined by the Infiniband
	// Architecture Specification.
	RoCEv2Protocol Protocol = "RoCEv2"
	// I2CProtocol shall mean that this device conforms to the NXP
	// Semiconductors I2C-bus Specification.
	I2CProtocol Protocol = "I2C"
	// OEMProtocol shall mean that this device conforms to an OEM specific
	// architecture and additional information may be included in the OEM
	// section.
	OEMProtocol Protocol = "OEM"
)
