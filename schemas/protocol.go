//
// SPDX-License-Identifier: BSD-3-Clause
//

// http://redfish.dmtf.org/schemas/v1/Protocol.json
// 2017.3 - #Protocol

package schemas

type Protocol string

const (
	// PCIeProtocol shall indicate conformance to the PCI-SIG PCI Express Base
	// Specification.
	PCIeProtocol Protocol = "PCIe"
	// AHCIProtocol shall indicate conformance to the Intel Advanced Host
	// Controller Interface (AHCI) Specification.
	AHCIProtocol Protocol = "AHCI"
	// UHCIProtocol shall indicate conformance to the Intel Universal Host
	// Controller Interface (UHCI) Specification, Enhanced Host Controller
	// Interface Specification, or the Extensible Host Controller Interface
	// Specification.
	UHCIProtocol Protocol = "UHCI"
	// SASProtocol shall indicate conformance to the T10 SAS Protocol Layer
	// Specification.
	SASProtocol Protocol = "SAS"
	// SATAProtocol shall indicate conformance to the Serial ATA International
	// Organization Serial ATA Specification.
	SATAProtocol Protocol = "SATA"
	// USBProtocol shall indicate conformance to the USB Implementers Forum
	// Universal Serial Bus Specification.
	USBProtocol Protocol = "USB"
	// NVMeProtocol shall indicate conformance to the Non-Volatile Memory Host
	// Controller Interface Specification.
	NVMeProtocol Protocol = "NVMe"
	// FCProtocol shall indicate conformance to the T11 Fibre Channel Physical and
	// Signaling Interface Specification.
	FCProtocol Protocol = "FC"
	// iSCSIProtocol shall indicate conformance to the IETF Internet Small Computer
	// Systems Interface (iSCSI) Specification.
	ISCSIProtocol Protocol = "iSCSI"
	// FCoEProtocol shall indicate conformance to the T11 FC-BB-5 Specification.
	FCoEProtocol Protocol = "FCoE"
	// FCPProtocol shall indicate conformance to the INCITS 481: Information
	// Technology - Fibre Channel Protocol for SCSI.
	FCPProtocol Protocol = "FCP"
	// FICONProtocol shall indicate conformance to the ANSI FC-SB-3 Single-Byte
	// Command Code Sets-3 Mapping Protocol for the Fibre Channel (FC) protocol.
	// Fibre Connection (FICON) is the IBM-proprietary name for this protocol.
	FICONProtocol Protocol = "FICON"
	// NVMeOverFabricsProtocol shall indicate conformance to the NVM Express over
	// Fabrics Specification.
	NVMeOverFabricsProtocol Protocol = "NVMeOverFabrics"
	// SMBProtocol shall indicate conformance to the Server Message Block (SMB), or
	// Common Internet File System (CIFS), protocol.
	SMBProtocol Protocol = "SMB"
	// NFSv3Protocol shall indicate conformance to the RFC1813-defined Network File
	// System (NFS) protocol.
	NFSv3Protocol Protocol = "NFSv3"
	// NFSv4Protocol Network File System (NFS) version 4.
	NFSv4Protocol Protocol = "NFSv4"
	// HTTPProtocol shall indicate conformance to the Hypertext Transport Protocol
	// (HTTP) as defined by RFC3010 or RFC5661.
	HTTPProtocol Protocol = "HTTP"
	// HTTPSProtocol shall indicate conformance to the Hypertext Transfer Protocol
	// Secure (HTTPS) as defined by RFC2068 or RFC2616, which uses Transport Layer
	// Security (TLS) as defined by RFC5246 or RFC6176.
	HTTPSProtocol Protocol = "HTTPS"
	// FTPProtocol shall indicate conformance to the RFC114-defined File Transfer
	// Protocol (FTP).
	FTPProtocol Protocol = "FTP"
	// SFTPProtocol shall indicate conformance to the RFC114-defined SSH File
	// Transfer Protocol (SFTP) that uses Transport Layer Security (TLS) as defined
	// by RFC5246 or RFC6176.
	SFTPProtocol Protocol = "SFTP"
	// iWARPProtocol shall indicate conformance to the RFC5042-defined Internet
	// Wide Area RDMA Protocol (iWARP) that uses the transport layer mechanisms as
	// defined by RFC5043 or RFC5044.
	IWARPProtocol Protocol = "iWARP"
	// RoCEProtocol shall indicate conformance to the InfiniBand Architecture
	// Specification-defined RDMA over Converged Ethernet Protocol.
	RoCEProtocol Protocol = "RoCE"
	// RoCEv2Protocol shall indicate conformance to the InfiniBand Architecture
	// Specification-defined RDMA over Converged Ethernet Protocol version 2.
	RoCEv2Protocol Protocol = "RoCEv2"
	// I2CProtocol shall indicate conformance to the NXP Semiconductors I2C-bus
	// Specification.
	I2CProtocol Protocol = "I2C"
	// TCPProtocol shall indicate conformance to the IETF-defined Transmission
	// Control Protocol (TCP). For example, RFC7414 defines the roadmap of the TCP
	// specification.
	TCPProtocol Protocol = "TCP"
	// UDPProtocol shall indicate conformance to the IETF-defined User Datagram
	// Protocol (UDP). For example, RFC768 defines the core UDP specification.
	UDPProtocol Protocol = "UDP"
	// TFTPProtocol shall indicate conformance to the IETF-defined Trivial File
	// Transfer Protocol (TFTP). For example, RFC1350 defines the core TFTP version
	// 2 specification.
	TFTPProtocol Protocol = "TFTP"
	// GenZProtocol shall indicate conformance to the Gen-Z Core Specification.
	GenZProtocol Protocol = "GenZ"
	// MultiProtocolProtocol shall indicate conformance to multiple protocols.
	MultiProtocolProtocol Protocol = "MultiProtocol"
	// InfiniBandProtocol shall indicate conformance to the InfiniBand Architecture
	// Specification-defined InfiniBand protocol.
	InfiniBandProtocol Protocol = "InfiniBand"
	// EthernetProtocol shall indicate conformance to the IEEE 802.3 Ethernet
	// specification.
	EthernetProtocol Protocol = "Ethernet"
	// NVLinkProtocol shall indicate conformance to the NVIDIA NVLink protocol.
	NVLinkProtocol Protocol = "NVLink"
	// OEMProtocol shall indicate conformance to an OEM-specific architecture, and
	// the OEM section may include additional information.
	OEMProtocol Protocol = "OEM"
	// DisplayPortProtocol shall indicate conformance to the VESA DisplayPort
	// Specification.
	DisplayPortProtocol Protocol = "DisplayPort"
	// HDMIProtocol shall indicate conformance to the HDMI Forum HDMI
	// Specification.
	HDMIProtocol Protocol = "HDMI"
	// VGAProtocol shall indicate conformance to the VESA SVGA Specification.
	VGAProtocol Protocol = "VGA"
	// DVIProtocol shall indicate conformance to the Digital Display Working Group
	// DVI-A, DVI-D, or DVI-I Specification.
	DVIProtocol Protocol = "DVI"
	// CXLProtocol shall indicate conformance to the Compute Express Link
	// Specification.
	CXLProtocol Protocol = "CXL"
	// UPIProtocol shall indicate conformance to the Intel UltraPath Interconnect
	// (UPI) protocol.
	UPIProtocol Protocol = "UPI"
	// QPIProtocol shall indicate conformance to the Intel QuickPath Interconnect
	// (QPI) protocol.
	QPIProtocol Protocol = "QPI"
	// eMMCProtocol shall indicate conformance to the JEDEC JESD84-B51A
	// specification.
	EMMCProtocol Protocol = "eMMC"
	// UETProtocol shall indicate conformance to the Ultra Ethernet specifications.
	UETProtocol Protocol = "UET"
	// UALinkProtocol shall indicate conformance to the Ultra Accelerator Link
	// specifications.
	UALinkProtocol Protocol = "UALink"
)
