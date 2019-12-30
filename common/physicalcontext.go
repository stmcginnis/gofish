//
// SPDX-License-Identifier: BSD-3-Clause
//

package common

// PhysicalContext is the physical location.
type PhysicalContext string

const (
	// RoomPhysicalContext The room.
	RoomPhysicalContext PhysicalContext = "Room"
	// IntakePhysicalContext The air intake point(s) or region of the chassis.
	IntakePhysicalContext PhysicalContext = "Intake"
	// ExhaustPhysicalContext The air exhaust point(s) or region of the chassis.
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
	// CPUPhysicalContext A Processor (CPU).
	CPUPhysicalContext PhysicalContext = "CPU"
	// CPUSubsystemPhysicalContext The entire Processor (CPU) subsystem.
	CPUSubsystemPhysicalContext PhysicalContext = "CPUSubsystem"
	// GPUPhysicalContext A Graphics Processor (GPU).
	GPUPhysicalContext PhysicalContext = "GPU"
	// GPUSubsystemPhysicalContext The entire Graphics Processor (GPU)
	// subsystem.
	GPUSubsystemPhysicalContext PhysicalContext = "GPUSubsystem"
	// FPGAPhysicalContext A Field Programmable Gate Array (FPGA).
	FPGAPhysicalContext PhysicalContext = "FPGA"
	// AcceleratorPhysicalContext An Accelerator.
	AcceleratorPhysicalContext PhysicalContext = "Accelerator"
	// ASICPhysicalContext An ASIC device, such as networking chip or a chipset
	// component.
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
	// MemorySubsystemPhysicalContext The entire Memory subsystem.
	MemorySubsystemPhysicalContext PhysicalContext = "MemorySubsystem"
	// ChassisPhysicalContext The entire chassis.
	ChassisPhysicalContext PhysicalContext = "Chassis"
	// FanPhysicalContext A fan.
	FanPhysicalContext PhysicalContext = "Fan"
	// CoolingSubsystemPhysicalContext The entire cooling (air and liquid)
	// subsystem.
	CoolingSubsystemPhysicalContext PhysicalContext = "CoolingSubsystem"
	// MotorPhysicalContext A motor.
	MotorPhysicalContext PhysicalContext = "Motor"
	// TransformerPhysicalContext A Transformer.
	TransformerPhysicalContext PhysicalContext = "Transformer"
	// ACUtilityInputPhysicalContext An AC Utility Input.
	ACUtilityInputPhysicalContext PhysicalContext = "ACUtilityInput"
	// ACStaticBypassInputPhysicalContext An AC Static Bypass Input.
	ACStaticBypassInputPhysicalContext PhysicalContext = "ACStaticBypassInput"
	// ACMaintenanceBypassInputPhysicalContext An AC Maintenance Bypass Input.
	ACMaintenanceBypassInputPhysicalContext PhysicalContext = "ACMaintenanceBypassInput"
	// DCBusPhysicalContext A DC Bus.
	DCBusPhysicalContext PhysicalContext = "DCBus"
	// ACOutputPhysicalContext An AC Output.
	ACOutputPhysicalContext PhysicalContext = "ACOutput"
	// ACInputPhysicalContext An AC Input.
	ACInputPhysicalContext PhysicalContext = "ACInput"
)

// PhysicalSubContext is the physical subcontext.
type PhysicalSubContext string

const (

	// InputPhysicalSubContext The input.
	InputPhysicalSubContext PhysicalSubContext = "Input"
	// OutputPhysicalSubContext The output.
	OutputPhysicalSubContext PhysicalSubContext = "Output"
)
