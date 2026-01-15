//
// SPDX-License-Identifier: BSD-3-Clause
//
// 2017.3 - #Protocol

package common

type Protocol string

const (
	// AHCIProtocol shall indicate conformance to the Intel Advanced Host
	// Controller Interface (AHCI) Specification.
	AHCIProtocol Protocol = "AHCI"
	// CXLProtocol shall indicate conformance to the Compute Express Link
	// Specification.
	CXLProtocol Protocol = "CXL"
	// DVIProtocol shall indicate conformance to the Digital Display Working
	// Group DVI-A, DVI-D, or DVI-I Specification.
	DVIProtocol Protocol = "DVI"
	// DisplayPortProtocol shall indicate conformance to the VESA DisplayPort
	// Specification.
	DisplayPortProtocol Protocol = "DisplayPort"
	// EthernetProtocol shall indicate conformance to the IEEE 802.3 Ethernet
	// specification.
	EthernetProtocol Protocol = "Ethernet"
	// FCProtocol shall indicate conformance to the T11 Fibre Channel Physical
	// and Signaling Interface Specification.
	FCProtocol Protocol = "FC"
	// FCPProtocol shall indicate conformance to the INCITS 481: Information
	// Technology - Fibre Channel Protocol for SCSI.
	FCPProtocol Protocol = "FCP"
	// FCoEProtocol shall indicate conformance to the T11 FC-BB-5 Specification.
	FCoEProtocol Protocol = "FCoE"
	// FICONProtocol shall indicate conformance to the ANSI FC-SB-3 Single-Byte
	// Command Code Sets-3 Mapping Protocol for the Fibre Channel (FC) protocol.
	// Fibre Connection (FICON) is the IBM-proprietary name for this protocol.
	FICONProtocol Protocol = "FICON"
	// FTPProtocol shall indicate conformance to the RFC114-defined File Transfer
	// Protocol (FTP).
	FTPProtocol Protocol = "FTP"
	// GenZProtocol shall indicate conformance to the Gen-Z Core Specification.
	GenZProtocol Protocol = "GenZ"
	// HDMIProtocol shall indicate conformance to the HDMI Forum HDMI Specification.
	HDMIProtocol Protocol = "HDMI"
	// HTTPProtocol shall indicate conformance to the Hypertext Transport Protocol
	// (HTTP) as defined by RFC3010 or RFC5661.
	HTTPProtocol Protocol = "HTTP"
	// HTTPSProtocol shall indicate conformance to the Hypertext Transfer Protocol
	// Secure (HTTPS) as defined by RFC2068 or RFC2616, which uses Transport Layer
	// Security (TLS) as defined by RFC5246 or RFC6176.
	HTTPSProtocol Protocol = "HTTPS"
	// I2CProtocol shall indicate conformance to the NXP Semiconductors I2C-bus
	// Specification.
	I2CProtocol Protocol = "I2C"
	// InfiniBandProtocol shall indicate conformance to the InfiniBand Architecture
	// Specification-defined InfiniBand protocol.
	InfiniBandProtocol Protocol = "InfiniBand"
	// MultiProtocolProtocol shall indicate conformance to multiple protocols.
	MultiProtocolProtocol Protocol = "MultiProtocol"
	// NFSv3Protocol shall indicate conformance to the RFC1813-defined Network
	// File System (NFS) protocol.
	NFSv3Protocol Protocol = "NFSv3"
	// NVLinkProtocol shall indicate conformance to the NVIDIA NVLink protocol.
	NVLinkProtocol Protocol = "NVLink"
	// NVMeProtocol shall indicate conformance to the Non-Volatile Memory Host
	// Controller Interface Specification.
	NVMeProtocol Protocol = "NVMe"
	// NVMeOverFabricsProtocol shall indicate conformance to the NVM Express over
	// Fabrics Specification.
	NVMeOverFabricsProtocol Protocol = "NVMeOverFabrics"
	// OEMProtocol shall indicate conformance to an OEM-specific architecture,
	// and the OEM section may include additional information.
	OEMProtocol Protocol = "OEM"
	// PCIeProtocol shall indicate conformance to the PCI-SIG PCI Express Base
	// Specification.
	PCIeProtocol Protocol = "PCIe"
	// QPIProtocol shall indicate conformance to the Intel QuickPath Interconnect
	// (QPI) protocol.
	QPIProtocol Protocol = "QPI"
	// RoCEProtocol shall indicate conformance to the InfiniBand Architecture
	// Specification-defined RDMA over Converged Ethernet Protocol.
	RoCEProtocol Protocol = "RoCE"
	// RoCEv2Protocol shall indicate conformance to the InfiniBand Architecture
	// Specification-defined RDMA over Converged Ethernet Protocol version 2.
	RoCEv2Protocol Protocol = "RoCEv2"
	// SASProtocol shall indicate conformance to the T10 SAS Protocol Layer
	// Specification.
	SASProtocol Protocol = "SAS"
	// SATAProtocol shall indicate conformance to the Serial ATA International
	// Organization Serial ATA Specification.
	SATAProtocol Protocol = "SATA"
	// SFTPProtocol shall indicate conformance to the RFC114-defined SSH File
	// Transfer Protocol (SFTP) that uses Transport Layer Security (TLS) as
	// defined by RFC5246 or RFC6176.
	SFTPProtocol Protocol = "SFTP"
	// SMBProtocol shall indicate conformance to the Server Message Block (SMB),
	// or Common Internet File System (CIFS), protocol.
	SMBProtocol Protocol = "SMB"
	// TCPProtocol shall indicate conformance to the IETF-defined Transmission
	// Control Protocol (TCP).  For example, RFC7414 defines the roadmap of the
	// TCP specification.
	TCPProtocol Protocol = "TCP"
	// TFTPProtocol shall indicate conformance to the IETF-defined Trivial File
	// Transfer Protocol (TFTP).  For example, RFC1350 defines the core TFTP
	// version 2 specification.
	TFTPProtocol Protocol = "TFTP"
	// UDPProtocol shall indicate conformance to the IETF-defined User Datagram
	// Protocol (UDP).  For example, RFC768 defines the core UDP specification.
	UDPProtocol Protocol = "UDP"
	// UETProtocol shall indicate conformance to the Ultra Ethernet specifications.
	UETProtocol Protocol = "UET"
	// UHCIProtocol shall indicate conformance to the Intel Universal Host
	// Controller Interface (UHCI) Specification, Enhanced Host Controller
	// Interface Specification, or the Extensible Host Controller Interface
	// Specification.
	UHCIProtocol Protocol = "UHCI"
	// UPIProtocol shall indicate conformance to the Intel UltraPath Interconnect
	// (UPI) protocol.
	UPIProtocol Protocol = "UPI"
	// USBProtocol shall indicate conformance to the USB Implementers Forum
	// Universal Serial Bus Specification.
	USBProtocol Protocol = "USB"
	// VGAProtocol shall indicate conformance to the VESA SVGA Specification.
	VGAProtocol Protocol = "VGA"
	// eMMCProtocol shall indicate conformance to the JEDEC JESD84-B51A specification.
	EMMCProtocol Protocol = "eMMC"
	// iSCSIProtocol shall indicate conformance to the IETF Internet Small
	// Computer Systems Interface (iSCSI) Specification.
	ISCSIProtocol Protocol = "iSCSI"
	// iWARPProtocol shall indicate conformance to the RFC5042-defined Internet
	// Wide Area RDMA Protocol (iWARP) that uses the transport layer mechanisms
	// as defined by RFC5043 or RFC5044."
	IWARPProtocol Protocol = "iWARP"
)
