//
// SPDX-License-Identifier: BSD-3-Clause
//
// 1.0 - #PhysicalContext

package schemas

type LogicalContext string

const (
	// CapacityLogicalContext Capacity-related logical context.
	CapacityLogicalContext LogicalContext = "Capacity"
	// EnvironmentLogicalContext Environment-related logical context.
	EnvironmentLogicalContext LogicalContext = "Environment"
	// NetworkLogicalContext Network-related logical context.
	NetworkLogicalContext LogicalContext = "Network"
	// PerformanceLogicalContext Performance-related logical context.
	PerformanceLogicalContext LogicalContext = "Performance"
	// SecurityLogicalContext Security-related logical context.
	SecurityLogicalContext LogicalContext = "Security"
	// StorageLogicalContext Storage-related logical context.
	StorageLogicalContext LogicalContext = "Storage"
)

type PhysicalContext string

const (
	// RoomPhysicalContext The room.
	RoomPhysicalContext PhysicalContext = "Room"
	// IntakePhysicalContext The air intake point or points or region of the
	// chassis.
	IntakePhysicalContext PhysicalContext = "Intake"
	// ExhaustPhysicalContext The air exhaust point or points or region of the
	// chassis.
	ExhaustPhysicalContext PhysicalContext = "Exhaust"
	// LiquidInletPhysicalContext The liquid inlet point of the chassis.
	LiquidInletPhysicalContext PhysicalContext = "LiquidInlet"
	// LiquidOutletPhysicalContext The liquid outlet point of the chassis.
	LiquidOutletPhysicalContext PhysicalContext = "LiquidOutlet"
	// FrontPhysicalContext The front of the chassis.
	FrontPhysicalContext PhysicalContext = "Front"
	// BackPhysicalContext The back of the chassis.
	BackPhysicalContext PhysicalContext = "Back"
	// UpperPhysicalContext The upper portion of the chassis.
	UpperPhysicalContext PhysicalContext = "Upper"
	// LowerPhysicalContext The lower portion of the chassis.
	LowerPhysicalContext PhysicalContext = "Lower"
	// CPUPhysicalContext is a processor (CPU).
	CPUPhysicalContext PhysicalContext = "CPU"
	// CPUSubsystemPhysicalContext The entire processor (CPU) subsystem.
	CPUSubsystemPhysicalContext PhysicalContext = "CPUSubsystem"
	// GPUPhysicalContext is a graphics processor (GPU).
	GPUPhysicalContext PhysicalContext = "GPU"
	// GPUSubsystemPhysicalContext The entire graphics processor (GPU) subsystem.
	GPUSubsystemPhysicalContext PhysicalContext = "GPUSubsystem"
	// FPGAPhysicalContext is an FPGA.
	FPGAPhysicalContext PhysicalContext = "FPGA"
	// AcceleratorPhysicalContext is an accelerator.
	AcceleratorPhysicalContext PhysicalContext = "Accelerator"
	// ASICPhysicalContext is an ASIC device, such as a networking chip or chipset
	// component.
	ASICPhysicalContext PhysicalContext = "ASIC"
	// BackplanePhysicalContext is a backplane within the chassis.
	BackplanePhysicalContext PhysicalContext = "Backplane"
	// SystemBoardPhysicalContext The system board (PCB).
	SystemBoardPhysicalContext PhysicalContext = "SystemBoard"
	// PowerSupplyPhysicalContext is a power supply.
	PowerSupplyPhysicalContext PhysicalContext = "PowerSupply"
	// PowerSubsystemPhysicalContext The entire power subsystem.
	PowerSubsystemPhysicalContext PhysicalContext = "PowerSubsystem"
	// VoltageRegulatorPhysicalContext is a voltage regulator device.
	VoltageRegulatorPhysicalContext PhysicalContext = "VoltageRegulator"
	// RectifierPhysicalContext is a rectifier device.
	RectifierPhysicalContext PhysicalContext = "Rectifier"
	// StorageDevicePhysicalContext is a storage device.
	StorageDevicePhysicalContext PhysicalContext = "StorageDevice"
	// StorageSubsystemPhysicalContext shall indicate a storage subsystem, which
	// may consist of one or more storage controllers, storage devices, or related
	// components.
	StorageSubsystemPhysicalContext PhysicalContext = "StorageSubsystem"
	// NetworkingDevicePhysicalContext is a networking device.
	NetworkingDevicePhysicalContext PhysicalContext = "NetworkingDevice"
	// ExpansionSubsystemPhysicalContext shall indicate a group of expansion bays
	// and the devices installed in those bays.
	ExpansionSubsystemPhysicalContext PhysicalContext = "ExpansionSubsystem"
	// ComputeBayPhysicalContext is a compute bay.
	ComputeBayPhysicalContext PhysicalContext = "ComputeBay"
	// StorageBayPhysicalContext shall indicate a location that provides for the
	// expansion of storage functionality of a system, by the addition of storage
	// devices.
	StorageBayPhysicalContext PhysicalContext = "StorageBay"
	// NetworkBayPhysicalContext shall indicate a location that provides for the
	// expansion of networking functionality of a system, by the addition of
	// networking devices.
	NetworkBayPhysicalContext PhysicalContext = "NetworkBay"
	// ExpansionBayPhysicalContext shall indicate a location that provides for the
	// expansion of functionality of a system, such as a PCIe slot that can accept
	// an option card.
	ExpansionBayPhysicalContext PhysicalContext = "ExpansionBay"
	// PowerSupplyBayPhysicalContext shall indicate a location that provides for
	// the installation of a power supply or similar devices.
	PowerSupplyBayPhysicalContext PhysicalContext = "PowerSupplyBay"
	// MemoryPhysicalContext is a memory device.
	MemoryPhysicalContext PhysicalContext = "Memory"
	// MemorySubsystemPhysicalContext The entire memory subsystem.
	MemorySubsystemPhysicalContext PhysicalContext = "MemorySubsystem"
	// ChassisPhysicalContext The entire chassis.
	ChassisPhysicalContext PhysicalContext = "Chassis"
	// FanPhysicalContext is a fan.
	FanPhysicalContext PhysicalContext = "Fan"
	// CoolingSubsystemPhysicalContext The entire cooling, or air and liquid,
	// subsystem.
	CoolingSubsystemPhysicalContext PhysicalContext = "CoolingSubsystem"
	// MotorPhysicalContext is a motor.
	MotorPhysicalContext PhysicalContext = "Motor"
	// TransformerPhysicalContext is a transformer.
	TransformerPhysicalContext PhysicalContext = "Transformer"
	// ACUtilityInputPhysicalContext shall indicate an electrical input, where the
	// source is an electrical utility as opposed to a backup or locally-generated
	// power source. This value is intended to differentiate multiple electrical
	// inputs between utility, maintenance bypass, or static bypass values. For
	// general purpose usage, the value of 'ACInput' is preferred.
	ACUtilityInputPhysicalContext PhysicalContext = "ACUtilityInput"
	// ACStaticBypassInputPhysicalContext is an AC electrical static bypass input.
	ACStaticBypassInputPhysicalContext PhysicalContext = "ACStaticBypassInput"
	// ACMaintenanceBypassInputPhysicalContext is an AC electrical maintenance
	// bypass input.
	ACMaintenanceBypassInputPhysicalContext PhysicalContext = "ACMaintenanceBypassInput"
	// DCBusPhysicalContext is a DC electrical bus.
	DCBusPhysicalContext PhysicalContext = "DCBus"
	// ACOutputPhysicalContext shall indicate an electrical output or an
	// output-related circuit, such as a branch output, which is not terminated as
	// a power outlet.
	ACOutputPhysicalContext PhysicalContext = "ACOutput"
	// ACInputPhysicalContext is an AC electrical input or input-related circuit.
	ACInputPhysicalContext PhysicalContext = "ACInput"
	// PowerOutletPhysicalContext shall indicate an electrical outlet or
	// receptacle.
	PowerOutletPhysicalContext PhysicalContext = "PowerOutlet"
	// TrustedModulePhysicalContext is a trusted module.
	TrustedModulePhysicalContext PhysicalContext = "TrustedModule"
	// BoardPhysicalContext shall indicate a circuit board that is not the primary
	// or system board within a context that cannot be described by other defined
	// values.
	BoardPhysicalContext PhysicalContext = "Board"
	// TransceiverPhysicalContext shall indicate a transceiver attached to a
	// device.
	TransceiverPhysicalContext PhysicalContext = "Transceiver"
	// BatteryPhysicalContext is a battery.
	BatteryPhysicalContext PhysicalContext = "Battery"
	// PumpPhysicalContext is a pump.
	PumpPhysicalContext PhysicalContext = "Pump"
	// FilterPhysicalContext is a filter.
	FilterPhysicalContext PhysicalContext = "Filter"
	// ReservoirPhysicalContext is a reservoir.
	ReservoirPhysicalContext PhysicalContext = "Reservoir"
	// SwitchPhysicalContext is a switch device.
	SwitchPhysicalContext PhysicalContext = "Switch"
	// ManagerPhysicalContext is a management controller, such as a BMC (baseboard
	// management controller).
	ManagerPhysicalContext PhysicalContext = "Manager"
)

type PhysicalSubContext string

const (
	// InputPhysicalSubContext The input.
	InputPhysicalSubContext PhysicalSubContext = "Input"
	// OutputPhysicalSubContext The output.
	OutputPhysicalSubContext PhysicalSubContext = "Output"
)
