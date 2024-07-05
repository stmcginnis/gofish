//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"

	"github.com/stmcginnis/gofish/common"
)

type ContainerEngineTypes string

const (
	// DockerContainerEngineTypes shall indicate the container engine is Docker.
	DockerContainerEngineTypes ContainerEngineTypes = "Docker"
	// containerdContainerEngineTypes shall indicate the container engine is containerd.
	ContainerdContainerEngineTypes ContainerEngineTypes = "containerd"
	// CRIOContainerEngineTypes shall indicate the container engine is CRI-O.
	CRIOContainerEngineTypes ContainerEngineTypes = "CRIO"
)

type OperatingSystemTypes string

const (
	// LinuxOperatingSystemTypes Linux.
	LinuxOperatingSystemTypes OperatingSystemTypes = "Linux"
	// WindowsOperatingSystemTypes Microsoft Windows.
	WindowsOperatingSystemTypes OperatingSystemTypes = "Windows"
	// SolarisOperatingSystemTypes Oracle Solaris.
	SolarisOperatingSystemTypes OperatingSystemTypes = "Solaris"
	// HPUXOperatingSystemTypes HPE HP-UX.
	HPUXOperatingSystemTypes OperatingSystemTypes = "HPUX"
	// AIXOperatingSystemTypes IBM AIX.
	AIXOperatingSystemTypes OperatingSystemTypes = "AIX"
	// BSDOperatingSystemTypes Berkeley Software Distribution.
	BSDOperatingSystemTypes OperatingSystemTypes = "BSD"
	// macOSOperatingSystemTypes Apple macOS.
	MacOSOperatingSystemTypes OperatingSystemTypes = "macOS"
	// IBMiOperatingSystemTypes IBM i.
	IBMiOperatingSystemTypes OperatingSystemTypes = "IBMi"
	// HypervisorOperatingSystemTypes A bare-metal hypervisor.
	HypervisorOperatingSystemTypes OperatingSystemTypes = "Hypervisor"
)

type VirtualMachineEngineTypes string

const (
	// VMwareESXVirtualMachineEngineTypes shall indicate the virtual machine engine is VMware ESX or ESXi.
	VMwareESXVirtualMachineEngineTypes VirtualMachineEngineTypes = "VMwareESX"
	// HyperVVirtualMachineEngineTypes shall indicate the virtual machine engine is Microsoft Hyper-V.
	HyperVVirtualMachineEngineTypes VirtualMachineEngineTypes = "HyperV"
	// XenVirtualMachineEngineTypes shall indicate the virtual machine engine is Xen.
	XenVirtualMachineEngineTypes VirtualMachineEngineTypes = "Xen"
	// KVMVirtualMachineEngineTypes shall indicate the virtual machine engine is Linux KVM (Kernel-based Virtual
	// Machine).
	KVMVirtualMachineEngineTypes VirtualMachineEngineTypes = "KVM"
	// QEMUVirtualMachineEngineTypes shall indicate the virtual machine engine is QEMU (Quick Emulator). If QEMU is
	// acting as a frontend for another virtual machine engine, such as Xen or KVM, VirtualMachineEngines should
	// contain additional entries to represent the backend virtual machine engines.
	QEMUVirtualMachineEngineTypes VirtualMachineEngineTypes = "QEMU"
	// VirtualBoxVirtualMachineEngineTypes shall indicate the virtual machine engine is Oracle VM VirtualBox. If
	// VirtualBox is acting as a frontend for another virtual machine engine, such as HyperV, VirtualMachineEngines
	// should contain additional entries to represent the backend virtual machine engines.
	VirtualBoxVirtualMachineEngineTypes VirtualMachineEngineTypes = "VirtualBox"
	// PowerVMVirtualMachineEngineTypes shall indicate the virtual machine engine is IBM PowerVM.
	PowerVMVirtualMachineEngineTypes VirtualMachineEngineTypes = "PowerVM"
)

type VirtualMachineImageTypes string

const (
	// RawVirtualMachineImageTypes shall indicate a raw disk image.
	RawVirtualMachineImageTypes VirtualMachineImageTypes = "Raw"
	// OVFVirtualMachineImageTypes shall indicate a DSP0243-defined OVF (Open Virtualization Format) image.
	OVFVirtualMachineImageTypes VirtualMachineImageTypes = "OVF"
	// OVAVirtualMachineImageTypes shall indicate a DSP0243-defined OVA (Open Virtual Appliance) image.
	OVAVirtualMachineImageTypes VirtualMachineImageTypes = "OVA"
	// VHDVirtualMachineImageTypes shall indicate a Microsoft Open Specification Promise-defined VHD (Virtual Hard
	// Disk) image.
	VHDVirtualMachineImageTypes VirtualMachineImageTypes = "VHD"
	// VMDKVirtualMachineImageTypes shall indicate a VMware-defined VMDK (Virtual Machine Disk) image.
	VMDKVirtualMachineImageTypes VirtualMachineImageTypes = "VMDK"
	// VDIVirtualMachineImageTypes shall indicate an Oracle VM VirtualBox-defined VDI (Virtual Disk Image).
	VDIVirtualMachineImageTypes VirtualMachineImageTypes = "VDI"
	// QCOWVirtualMachineImageTypes shall indicate a QEMU-defined QCOW (QEMU Copy-on-Write) image.
	QCOWVirtualMachineImageTypes VirtualMachineImageTypes = "QCOW"
	// QCOW2VirtualMachineImageTypes shall indicate a QEMU-defined QCOW2 (QEMU Copy-on-Write version 2) image.
	QCOW2VirtualMachineImageTypes VirtualMachineImageTypes = "QCOW2"
)

// ContainerEngine shall contain a container engine running in an operating system.
type ContainerEngine struct {
	// ManagementURIs shall contain an array of URIs to management interfaces for this container engine. This is
	// typically a web UI or API provided by the container engine.
	ManagementURIs []string
	// SupportedImageTypes shall contain the supported image types for this container engine.
	SupportedImageTypes []ImageTypes
	// Type shall contain the type for this container engine.
	Type ContainerEngineTypes
	// Version shall contain the version of this container engine.
	Version string
}

// Kernel shall contain the kernel information for an operating system.
type Kernel struct {
	common.Entity
	// Machine shall contain the machine hardware name of the kernel. For strict POSIX operating systems, the value
	// shall contain the output of 'uname -m'.
	Machine string
	// Release shall contain the release of the kernel. For strict POSIX operating systems, the value shall contain the
	// output of 'uname -r'. For Microsoft Windows, the value shall contain the decimal-delimited version from the
	// output of 'ver', from Command Prompt, within the square braces ('[' and ']'), following the regular expression
	// '^\d+\.\d+\.\d+\.\d+$'.
	Release string
	// Version shall contain the version of the kernel. For strict POSIX operating systems, the value shall contain the
	// output of 'uname -v'.
	Version string
}

// OperatingSystem shall represent the operating system and software running on a computer system.
type OperatingSystem struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Applications shall contain a link to a resource collection of type ApplicationCollection that represent the
	// applications running under this operating system.
	applications string
	// ContainerEngines shall contain the container engines running in this operating system.
	ContainerEngines []ContainerEngine
	// ContainerImages shall contain a link to a resource collection of type ContainerImageCollection that represent
	// the container images available to container engines on this operating system.
	containerImages string
	// Containers shall contain a link to a resource collection of type ContainerCollection that represent the
	// containers running under this operating system.
	containers string
	// Description provides a description of this resource.
	Description string
	// Kernel shall contain the kernel information for this operating system.
	Kernel Kernel
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// Status shall contain any status or health properties of the resource.
	Status common.Status
	// Type shall contain the type for this operating system.
	Type OperatingSystemTypes
	// UptimeSeconds shall contain the wall-clock time this operating system has been running in seconds.
	UptimeSeconds int
	// VirtualMachineEngines shall contain the virtual machine engines running in this operating system.
	VirtualMachineEngines []VirtualMachineEngine

	softwareImage string
}

// UnmarshalJSON unmarshals a OperatingSystem object from the raw JSON.
func (operatingsystem *OperatingSystem) UnmarshalJSON(b []byte) error {
	type temp OperatingSystem
	type Links struct {
		SoftwareImage common.Link
	}
	var t struct {
		temp
		Links           Links
		Applications    common.Link
		ContainerImages common.Link
		Containers      common.Link
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*operatingsystem = OperatingSystem(t.temp)

	// Extract the links to other entities for later
	operatingsystem.softwareImage = t.Links.SoftwareImage.String()

	operatingsystem.applications = t.Applications.String()
	operatingsystem.containerImages = t.ContainerImages.String()
	operatingsystem.containers = t.Containers.String()

	return nil
}

// SoftwareImage gets the software image from which this operating system runs.
func (operatingsystem *OperatingSystem) SoftwareImage() (*SoftwareInventory, error) {
	if operatingsystem.softwareImage == "" {
		return nil, nil
	}
	return GetSoftwareInventory(operatingsystem.GetClient(), operatingsystem.softwareImage)
}

// Applications gets the applications running under this operating system.
func (operatingsystem *OperatingSystem) Applications() ([]*Application, error) {
	return ListReferencedApplications(operatingsystem.GetClient(), operatingsystem.applications)
}

// ContainerImages gets the container images available to container engines on this operating system.
func (operatingsystem *OperatingSystem) ContainerImages() ([]*ContainerImage, error) {
	return ListReferencedContainerImages(operatingsystem.GetClient(), operatingsystem.containerImages)
}

// Containers gets the containers running under this operating system.
func (operatingsystem *OperatingSystem) Containers() ([]*Container, error) {
	return ListReferencedContainers(operatingsystem.GetClient(), operatingsystem.containers)
}

// GetOperatingSystem will get a OperatingSystem instance from the service.
func GetOperatingSystem(c common.Client, uri string) (*OperatingSystem, error) {
	return common.GetObject[OperatingSystem](c, uri)
}

// ListReferencedOperatingSystems gets the collection of OperatingSystem from
// a provided reference.
func ListReferencedOperatingSystems(c common.Client, link string) ([]*OperatingSystem, error) {
	return common.GetCollectionObjects[OperatingSystem](c, link)
}

// VirtualMachineEngine shall contain a virtual machine engine running in an operating system.
type VirtualMachineEngine struct {
	// ManagementURIs shall contain an array of URIs to management interfaces for this virtual machine engine. This is
	// typically a web UI or API provided by the virtual machine engine.
	ManagementURIs []string
	// SupportedImageTypes shall contain the supported image types for this container engine.
	SupportedImageTypes []VirtualMachineImageTypes
	// Type shall contain the type for this virtual machine engine.
	Type VirtualMachineEngineTypes
	// Version shall contain the version of this virtual machine engine.
	Version string
}

// UnmarshalJSON unmarshals a VirtualMachineEngine object from the raw JSON.
func (virtualmachineengine *VirtualMachineEngine) UnmarshalJSON(b []byte) error {
	type temp VirtualMachineEngine
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*virtualmachineengine = VirtualMachineEngine(t.temp)

	// Extract the links to other entities for later

	return nil
}
