//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

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
	// IntakePhysicalContext The air intake point or points or region of the chassis.
	IntakePhysicalContext PhysicalContext = "Intake"
	// ExhaustPhysicalContext The air exhaust point or points or region of the chassis.
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
	// CPUPhysicalContext A processor (CPU).
	CPUPhysicalContext PhysicalContext = "CPU"
	// CPUSubsystemPhysicalContext The entire processor (CPU) subsystem.
	CPUSubsystemPhysicalContext PhysicalContext = "CPUSubsystem"
	// GPUPhysicalContext A graphics processor (GPU).
	GPUPhysicalContext PhysicalContext = "GPU"
	// GPUSubsystemPhysicalContext The entire graphics processor (GPU) subsystem.
	GPUSubsystemPhysicalContext PhysicalContext = "GPUSubsystem"
	// FPGAPhysicalContext An FPGA.
	FPGAPhysicalContext PhysicalContext = "FPGA"
	// AcceleratorPhysicalContext An accelerator.
	AcceleratorPhysicalContext PhysicalContext = "Accelerator"
	// ASICPhysicalContext An ASIC device, such as a networking chip or chipset component.
	ASICPhysicalContext PhysicalContext = "ASIC"
	// BackplanePhysicalContext A backplane within the chassis.
	BackplanePhysicalContext PhysicalContext = "Backplane"
	// SystemBoardPhysicalContext The system board (PCB).
	SystemBoardPhysicalContext PhysicalContext = "SystemBoard"
	// PowerSupplyPhysicalContext A power supply.
	PowerSupplyPhysicalContext PhysicalContext = "PowerSupply"
	// PowerSubsystemPhysicalContext The entire power subsystem.
	PowerSubsystemPhysicalContext PhysicalContext = "PowerSubsystem"
	// VoltageRegulatorPhysicalContext A voltage regulator device.
	VoltageRegulatorPhysicalContext PhysicalContext = "VoltageRegulator"
	// RectifierPhysicalContext A rectifier device.
	RectifierPhysicalContext PhysicalContext = "Rectifier"
	// StorageDevicePhysicalContext A storage device.
	StorageDevicePhysicalContext PhysicalContext = "StorageDevice"
	// NetworkingDevicePhysicalContext A networking device.
	NetworkingDevicePhysicalContext PhysicalContext = "NetworkingDevice"
	// ComputeBayPhysicalContext Within a compute bay.
	ComputeBayPhysicalContext PhysicalContext = "ComputeBay"
	// StorageBayPhysicalContext Within a storage bay.
	StorageBayPhysicalContext PhysicalContext = "StorageBay"
	// NetworkBayPhysicalContext Within a networking bay.
	NetworkBayPhysicalContext PhysicalContext = "NetworkBay"
	// ExpansionBayPhysicalContext Within an expansion bay.
	ExpansionBayPhysicalContext PhysicalContext = "ExpansionBay"
	// PowerSupplyBayPhysicalContext Within a power supply bay.
	PowerSupplyBayPhysicalContext PhysicalContext = "PowerSupplyBay"
	// MemoryPhysicalContext A memory device.
	MemoryPhysicalContext PhysicalContext = "Memory"
	// MemorySubsystemPhysicalContext The entire memory subsystem.
	MemorySubsystemPhysicalContext PhysicalContext = "MemorySubsystem"
	// ChassisPhysicalContext The entire chassis.
	ChassisPhysicalContext PhysicalContext = "Chassis"
	// FanPhysicalContext A fan.
	FanPhysicalContext PhysicalContext = "Fan"
	// CoolingSubsystemPhysicalContext The entire cooling, or air and liquid, subsystem.
	CoolingSubsystemPhysicalContext PhysicalContext = "CoolingSubsystem"
	// MotorPhysicalContext A motor.
	MotorPhysicalContext PhysicalContext = "Motor"
	// TransformerPhysicalContext A transformer.
	TransformerPhysicalContext PhysicalContext = "Transformer"
	// ACUtilityInputPhysicalContext An AC utility input.
	ACUtilityInputPhysicalContext PhysicalContext = "ACUtilityInput"
	// ACStaticBypassInputPhysicalContext An AC static bypass input.
	ACStaticBypassInputPhysicalContext PhysicalContext = "ACStaticBypassInput"
	// ACMaintenanceBypassInputPhysicalContext An AC maintenance bypass input.
	ACMaintenanceBypassInputPhysicalContext PhysicalContext = "ACMaintenanceBypassInput"
	// DCBusPhysicalContext A DC bus.
	DCBusPhysicalContext PhysicalContext = "DCBus"
	// ACOutputPhysicalContext An AC output.
	ACOutputPhysicalContext PhysicalContext = "ACOutput"
	// ACInputPhysicalContext An AC input.
	ACInputPhysicalContext PhysicalContext = "ACInput"
	// TrustedModulePhysicalContext A trusted module.
	TrustedModulePhysicalContext PhysicalContext = "TrustedModule"
	// BoardPhysicalContext shall indicate a circuit board that is not the primary or system board within a context
	// that cannot be described by other defined values.
	BoardPhysicalContext PhysicalContext = "Board"
	// TransceiverPhysicalContext shall indicate a transceiver attached to a device.
	TransceiverPhysicalContext PhysicalContext = "Transceiver"
	// BatteryPhysicalContext A battery.
	BatteryPhysicalContext PhysicalContext = "Battery"
	// PumpPhysicalContext A pump.
	PumpPhysicalContext PhysicalContext = "Pump"
)

type PhysicalSubContext string

const (
	// InputPhysicalSubContext The input.
	InputPhysicalSubContext PhysicalSubContext = "Input"
	// OutputPhysicalSubContext The output.
	OutputPhysicalSubContext PhysicalSubContext = "Output"
)
