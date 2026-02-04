//
// SPDX-License-Identifier: BSD-3-Clause
//

// http://redfish.dmtf.org/schemas/v1/OperatingSystem.v1_0_2.json
// 2023.2 - #OperatingSystem.v1_0_2.OperatingSystem

package schemas

import (
	"encoding/json"
)

type ContainerEngineTypes string

const (
	// DockerContainerEngineTypes shall indicate the container engine is Docker.
	DockerContainerEngineTypes ContainerEngineTypes = "Docker"
	// containerdContainerEngineTypes shall indicate the container engine is
	// containerd.
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
	// HypervisorOperatingSystemTypes is a bare-metal hypervisor.
	HypervisorOperatingSystemTypes OperatingSystemTypes = "Hypervisor"
)

type VirtualMachineEngineTypes string

const (
	// VMwareESXVirtualMachineEngineTypes shall indicate the virtual machine engine
	// is VMware ESX or ESXi.
	VMwareESXVirtualMachineEngineTypes VirtualMachineEngineTypes = "VMwareESX"
	// HyperVVirtualMachineEngineTypes shall indicate the virtual machine engine is
	// Microsoft Hyper-V.
	HyperVVirtualMachineEngineTypes VirtualMachineEngineTypes = "HyperV"
	// XenVirtualMachineEngineTypes shall indicate the virtual machine engine is
	// Xen.
	XenVirtualMachineEngineTypes VirtualMachineEngineTypes = "Xen"
	// KVMVirtualMachineEngineTypes shall indicate the virtual machine engine is
	// Linux KVM (Kernel-based Virtual Machine).
	KVMVirtualMachineEngineTypes VirtualMachineEngineTypes = "KVM"
	// QEMUVirtualMachineEngineTypes shall indicate the virtual machine engine is
	// QEMU (Quick Emulator). If QEMU is acting as a frontend for another virtual
	// machine engine, such as Xen or KVM, VirtualMachineEngines should contain
	// additional entries to represent the backend virtual machine engines.
	QEMUVirtualMachineEngineTypes VirtualMachineEngineTypes = "QEMU"
	// VirtualBoxVirtualMachineEngineTypes shall indicate the virtual machine
	// engine is Oracle VM VirtualBox. If VirtualBox is acting as a frontend for
	// another virtual machine engine, such as HyperV, VirtualMachineEngines should
	// contain additional entries to represent the backend virtual machine engines.
	VirtualBoxVirtualMachineEngineTypes VirtualMachineEngineTypes = "VirtualBox"
	// PowerVMVirtualMachineEngineTypes shall indicate the virtual machine engine
	// is IBM PowerVM.
	PowerVMVirtualMachineEngineTypes VirtualMachineEngineTypes = "PowerVM"
)

type VirtualMachineImageTypes string

const (
	// RawVirtualMachineImageTypes shall indicate a raw disk image.
	RawVirtualMachineImageTypes VirtualMachineImageTypes = "Raw"
	// OVFVirtualMachineImageTypes shall indicate a DSP0243-defined OVF (Open
	// Virtualization Format) image.
	OVFVirtualMachineImageTypes VirtualMachineImageTypes = "OVF"
	// OVAVirtualMachineImageTypes shall indicate a DSP0243-defined OVA (Open
	// Virtual Appliance) image.
	OVAVirtualMachineImageTypes VirtualMachineImageTypes = "OVA"
	// VHDVirtualMachineImageTypes shall indicate a Microsoft Open Specification
	// Promise-defined VHD (Virtual Hard Disk) image.
	VHDVirtualMachineImageTypes VirtualMachineImageTypes = "VHD"
	// VMDKVirtualMachineImageTypes shall indicate a VMware-defined VMDK (Virtual
	// Machine Disk) image.
	VMDKVirtualMachineImageTypes VirtualMachineImageTypes = "VMDK"
	// VDIVirtualMachineImageTypes shall indicate an Oracle VM VirtualBox-defined
	// VDI (Virtual Disk Image).
	VDIVirtualMachineImageTypes VirtualMachineImageTypes = "VDI"
	// QCOWVirtualMachineImageTypes shall indicate a QEMU-defined QCOW (QEMU
	// Copy-on-Write) image.
	QCOWVirtualMachineImageTypes VirtualMachineImageTypes = "QCOW"
	// QCOW2VirtualMachineImageTypes shall indicate a QEMU-defined QCOW2 (QEMU
	// Copy-on-Write version 2) image.
	QCOW2VirtualMachineImageTypes VirtualMachineImageTypes = "QCOW2"
)

// OperatingSystem shall represent the operating system and software running on
// a computer system.
type OperatingSystem struct {
	Entity
	// Applications shall contain a link to a resource collection of type
	// 'ApplicationCollection' that represent the applications running under this
	// operating system.
	applications string
	// ContainerEngines shall contain the container engines running in this
	// operating system.
	ContainerEngines []ContainerEngine
	// ContainerImages shall contain a link to a resource collection of type
	// 'ContainerImageCollection' that represent the container images available to
	// container engines on this operating system.
	containerImages string
	// Containers shall contain a link to a resource collection of type
	// 'ContainerCollection' that represent the containers running under this
	// operating system.
	containers string
	// Kernel shall contain the kernel information for this operating system.
	Kernel Kernel
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// OEM shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// Status shall contain any status or health properties of the resource.
	Status Status
	// Type shall contain the type for this operating system.
	Type OperatingSystemTypes
	// UptimeSeconds shall contain the wall-clock time this operating system has
	// been running in seconds.
	UptimeSeconds *int `json:",omitempty"`
	// VirtualMachineEngines shall contain the virtual machine engines running in
	// this operating system.
	VirtualMachineEngines []VirtualMachineEngine
	// softwareImage is the URI for SoftwareImage.
	softwareImage string
}

// UnmarshalJSON unmarshals a OperatingSystem object from the raw JSON.
func (o *OperatingSystem) UnmarshalJSON(b []byte) error {
	type temp OperatingSystem
	type oLinks struct {
		SoftwareImage Link `json:"SoftwareImage"`
	}
	var tmp struct {
		temp
		Links           oLinks
		Applications    Link `json:"Applications"`
		ContainerImages Link `json:"ContainerImages"`
		Containers      Link `json:"Containers"`
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*o = OperatingSystem(tmp.temp)

	// Extract the links to other entities for later
	o.softwareImage = tmp.Links.SoftwareImage.String()
	o.applications = tmp.Applications.String()
	o.containerImages = tmp.ContainerImages.String()
	o.containers = tmp.Containers.String()

	return nil
}

// GetOperatingSystem will get a OperatingSystem instance from the service.
func GetOperatingSystem(c Client, uri string) (*OperatingSystem, error) {
	return GetObject[OperatingSystem](c, uri)
}

// ListReferencedOperatingSystems gets the collection of OperatingSystem from
// a provided reference.
func ListReferencedOperatingSystems(c Client, link string) ([]*OperatingSystem, error) {
	return GetCollectionObjects[OperatingSystem](c, link)
}

// SoftwareImage gets the SoftwareImage linked resource.
func (o *OperatingSystem) SoftwareImage() (*SoftwareInventory, error) {
	if o.softwareImage == "" {
		return nil, nil
	}
	return GetObject[SoftwareInventory](o.client, o.softwareImage)
}

// Applications gets the Applications collection.
func (o *OperatingSystem) Applications() ([]*Application, error) {
	if o.applications == "" {
		return nil, nil
	}
	return GetCollectionObjects[Application](o.client, o.applications)
}

// ContainerImages gets the ContainerImages collection.
func (o *OperatingSystem) ContainerImages() ([]*ContainerImage, error) {
	if o.containerImages == "" {
		return nil, nil
	}
	return GetCollectionObjects[ContainerImage](o.client, o.containerImages)
}

// Containers gets the Containers collection.
func (o *OperatingSystem) Containers() ([]*Container, error) {
	if o.containers == "" {
		return nil, nil
	}
	return GetCollectionObjects[Container](o.client, o.containers)
}

// ContainerEngine shall contain a container engine running in an operating
// system.
type ContainerEngine struct {
	// ManagementURIs shall contain an array of URIs to management interfaces for
	// this container engine. This is typically a web UI or API provided by the
	// container engine.
	ManagementURIs []string
	// SupportedImageTypes shall contain the supported image types for this
	// container engine.
	SupportedImageTypes []ImageTypes
	// Type shall contain the type for this container engine.
	Type ContainerEngineTypes
	// Version shall contain the version of this container engine.
	Version string
}

// Kernel shall contain the kernel information for an operating system.
type Kernel struct {
	// Machine shall contain the machine hardware name of the kernel. For strict
	// POSIX operating systems, the value shall contain the output of 'uname -m'.
	Machine string
	// Name is the name of the resource or array element.
	Name string
	// Release shall contain the release of the kernel. For strict POSIX operating
	// systems, the value shall contain the output of 'uname -r'. For Microsoft
	// Windows, the value shall contain the decimal-delimited version from the
	// output of 'ver', from Command Prompt, within the square braces ('[' and
	// ']'), following the regular expression '^\d+\.\d+\.\d+\.\d+$'.
	Release string
	// Version shall contain the version of the kernel. For strict POSIX operating
	// systems, the value shall contain the output of 'uname -v'.
	Version string
}

// VirtualMachineEngine shall contain a virtual machine engine running in an
// operating system.
type VirtualMachineEngine struct {
	// ManagementURIs shall contain an array of URIs to management interfaces for
	// this virtual machine engine. This is typically a web UI or API provided by
	// the virtual machine engine.
	ManagementURIs []string
	// SupportedImageTypes shall contain the supported image types for this
	// container engine.
	SupportedImageTypes []VirtualMachineImageTypes
	// Type shall contain the type for this virtual machine engine.
	Type VirtualMachineEngineTypes
	// Version shall contain the version of this virtual machine engine.
	Version string
}
