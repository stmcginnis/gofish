//
// SPDX-License-Identifier: BSD-3-Clause
//
// 2025.3 - #PortMetrics.v1_8_1.PortMetrics

package schemas

import (
	"encoding/json"
)

// PortMetrics shall represent the port metrics for a switch device or component
// port summary in a Redfish implementation.
type PortMetrics struct {
	Entity
	// CXL shall contain the port metrics specific to CXL ports.
	//
	// Version added: v1.4.0
	CXL PortMetricsCXL
	// FibreChannel shall contain Fibre Channel-specific port metrics for network
	// ports.
	//
	// Version added: v1.2.0
	FibreChannel PortMetricsFibreChannel
	// GenZ shall contain the port metrics specific to Gen-Z ports.
	GenZ GenZPortMetrics
	// Networking shall contain port metrics for network ports, including Ethernet,
	// Fibre Channel, and InfiniBand, that are not specific to one of these
	// protocols.
	//
	// Version added: v1.1.0
	Networking Networking
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// OEM shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// PCIeErrors shall contain the PCIe errors associated with this port.
	//
	// Version added: v1.3.0
	PCIeErrors PCIeErrors
	// PCIeMetrics shall contain the PCIe metrics associated with this port.
	//
	// Version added: v1.8.0
	PCIeMetrics PCIeMetrics
	// RXBytes shall contain the total number of bytes received on a port since
	// reset, including host and remote management passthrough traffic, and
	// inclusive of all protocol overhead.
	//
	// Version added: v1.1.0
	RXBytes *int `json:",omitempty"`
	// RXErrors shall contain the total number of received errors on a port since
	// reset.
	//
	// Version added: v1.1.0
	RXErrors *int `json:",omitempty"`
	// SAS shall contain an array of physical-related metrics for Serial Attached
	// SCSI (SAS). Each member in the array shall represent a single phy.
	//
	// Version added: v1.1.0
	SAS []SAS
	// TXBytes shall contain the total number of bytes transmitted on a port since
	// reset, including host and remote management passthrough traffic, and
	// inclusive of all protocol overhead.
	//
	// Version added: v1.1.0
	TXBytes *int `json:",omitempty"`
	// TXErrors shall contain the total number of transmission errors on a port
	// since reset.
	//
	// Version added: v1.1.0
	TXErrors *int `json:",omitempty"`
	// Transceivers shall contain an array of transceiver-related metrics for this
	// port. Each member in the array shall represent a single transceiver.
	//
	// Version added: v1.1.0
	Transceivers []Transceiver
	// resetMetricsTarget is the URL to send ResetMetrics requests.
	resetMetricsTarget string
}

// UnmarshalJSON unmarshals a PortMetrics object from the raw JSON.
func (p *PortMetrics) UnmarshalJSON(b []byte) error {
	type temp PortMetrics
	type pActions struct {
		ResetMetrics ActionTarget `json:"#PortMetrics.ResetMetrics"`
	}
	var tmp struct {
		temp
		Actions pActions
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*p = PortMetrics(tmp.temp)

	// Extract the links to other entities for later
	p.resetMetricsTarget = tmp.Actions.ResetMetrics.Target

	return nil
}

// GetPortMetrics will get a PortMetrics instance from the service.
func GetPortMetrics(c Client, uri string) (*PortMetrics, error) {
	return GetObject[PortMetrics](c, uri)
}

// ListReferencedPortMetricss gets the collection of PortMetrics from
// a provided reference.
func ListReferencedPortMetricss(c Client, link string) ([]*PortMetrics, error) {
	return GetCollectionObjects[PortMetrics](c, link)
}

// This action shall reset any time intervals or counted values for this
// device.
//
// If TaskMonitorInfo is not nil it can be used to monitor async tasks.
func (p *PortMetrics) ResetMetrics() (*TaskMonitorInfo, error) {
	payload := make(map[string]any)
	resp, taskInfo, err := PostWithTask(p.client,
		p.resetMetricsTarget, payload, p.Headers(), false)
	defer DeferredCleanupHTTPResponse(resp)
	return taskInfo, err
}

// PortMetricsCXL shall contain the port metrics specific to CXL ports.
type PortMetricsCXL struct {
	// BackpressureAveragePercentage shall contain CXL Specification-defined
	// 'Backpressure Average Percentage' as a percentage, typically '0' to '100'.
	//
	// Version added: v1.4.0
	BackpressureAveragePercentage int
}

// PortMetricsFibreChannel shall describe Fibre Channel-specific metrics for network ports.
type PortMetricsFibreChannel struct {
	// CorrectableFECErrors shall contain the total number of times this port has
	// received traffic with correctable forward error correction (FEC) errors.
	//
	// Version added: v1.2.0
	CorrectableFECErrors *int `json:",omitempty"`
	// InvalidCRCs shall contain the total number of invalid cyclic redundancy
	// checks (CRCs) observed on this port.
	//
	// Version added: v1.2.0
	InvalidCRCs *int `json:",omitempty"`
	// InvalidTXWords shall contain the total number of times this port has
	// received invalid transmission words.
	//
	// Version added: v1.2.0
	InvalidTXWords *int `json:",omitempty"`
	// LinkFailures shall contain the total number of link failures observed on
	// this port.
	//
	// Version added: v1.2.0
	LinkFailures *int `json:",omitempty"`
	// LossesOfSignal shall contain the total number of times this port has lost
	// signal.
	//
	// Version added: v1.2.0
	LossesOfSignal *int `json:",omitempty"`
	// LossesOfSync shall contain the total number of times this port has lost
	// sync.
	//
	// Version added: v1.2.0
	LossesOfSync *int `json:",omitempty"`
	// RXBBCreditZero shall contain the number of times the receive
	// buffer-to-buffer credit count transitioned to zero since last counter reset.
	//
	// Version added: v1.2.0
	RXBBCreditZero *int `json:",omitempty"`
	// RXExchanges shall contain the total number of Fibre Channel exchanges
	// received.
	//
	// Version added: v1.2.0
	RXExchanges *int `json:",omitempty"`
	// RXSequences shall contain the total number of Fibre Channel sequences
	// received.
	//
	// Version added: v1.2.0
	RXSequences *int `json:",omitempty"`
	// TXBBCreditZero shall contain the number of times the transmit
	// buffer-to-buffer credit count transitioned to zero since last counter reset.
	//
	// Version added: v1.2.0
	TXBBCreditZero *int `json:",omitempty"`
	// TXBBCreditZeroDurationMilliseconds shall contain the total amount of time in
	// milliseconds the port has been blocked from transmitting due to lack of
	// buffer credits since the last counter reset.
	//
	// Version added: v1.2.0
	TXBBCreditZeroDurationMilliseconds *int `json:",omitempty"`
	// TXBBCredits shall contain the number of transmit buffer-to-buffer credits
	// the port is configured to use.
	//
	// Version added: v1.2.0
	TXBBCredits *int `json:",omitempty"`
	// TXExchanges shall contain the total number of Fibre Channel exchanges
	// transmitted.
	//
	// Version added: v1.2.0
	TXExchanges *int `json:",omitempty"`
	// TXSequences shall contain the total number of Fibre Channel sequences
	// transmitted.
	//
	// Version added: v1.2.0
	TXSequences *int `json:",omitempty"`
	// UncorrectableFECErrors shall contain the total number of times this port has
	// received traffic with uncorrectable forward error correction (FEC) errors.
	//
	// Version added: v1.2.0
	UncorrectableFECErrors *int `json:",omitempty"`
}

// GenZPortMetrics shall describe the Gen-Z related port metrics.
type GenZPortMetrics struct {
	// AccessKeyViolations shall contain the total number of Access Key Violations
	// detected for packets received or transmitted on this interface.
	AccessKeyViolations *int `json:",omitempty"`
	// EndToEndCRCErrors shall contain total number of ECRC transient errors
	// detected in received link-local and end-to-end packets.
	EndToEndCRCErrors *int `json:",omitempty"`
	// LLRRecovery shall contain the total number of times Link-level Reliability
	// (LLR) recovery has been initiated by this interface. This is not to be
	// confused with the number of packets retransmitted due to initiating LLR
	// recovery.
	LLRRecovery *int `json:",omitempty"`
	// LinkNTE shall contain the total number of link-local non-transient errors
	// detected on this interface.
	LinkNTE *int `json:",omitempty"`
	// MarkedECN shall contain the number of packets that the component set the
	// Congestion ECN bit prior to transmission through this interface.
	MarkedECN *int `json:",omitempty"`
	// NonCRCTransientErrors shall contain the total number of transient errors
	// detected that are unrelated to CRC validation, which covers link-local and
	// end-to-end packets, such as malformed Link Idle packets or PLA signal
	// errors.
	NonCRCTransientErrors *int `json:",omitempty"`
	// PacketCRCErrors shall contain the total number of PCRC transient errors
	// detected in received link-local and end-to-end packets.
	PacketCRCErrors *int `json:",omitempty"`
	// PacketDeadlineDiscards shall contain the number of packets discarded by this
	// interface due to the Congestion Deadline subfield reaching zero prior to
	// packet transmission.
	PacketDeadlineDiscards *int `json:",omitempty"`
	// RXStompedECRC shall contain the total number of packets that this interface
	// received with a stomped ECRC field.
	RXStompedECRC *int `json:",omitempty"`
	// ReceivedECN shall contain the number of packets received on this interface
	// with the Congestion ECN bit set.
	ReceivedECN *int `json:",omitempty"`
	// TXStompedECRC shall contain the total number of packets that this interfaced
	// stomped the ECRC field.
	TXStompedECRC *int `json:",omitempty"`
}

// Networking shall describe the metrics for network ports, including Ethernet,
// Fibre Channel, and InfiniBand, that are not specific to one of these
// protocols.
type Networking struct {
	// RDMAProtectionErrors shall contain the total number of RDMA protection
	// errors.
	//
	// Version added: v1.1.0
	RDMAProtectionErrors *int `json:",omitempty"`
	// RDMAProtocolErrors shall contain the total number of RDMA protocol errors.
	//
	// Version added: v1.1.0
	RDMAProtocolErrors *int `json:",omitempty"`
	// RDMARXBytes shall contain the total number of RDMA bytes received on a port
	// since reset.
	//
	// Version added: v1.1.0
	RDMARXBytes *int `json:",omitempty"`
	// RDMARXRequests shall contain the total number of RDMA requests received on a
	// port since reset.
	//
	// Version added: v1.1.0
	RDMARXRequests *int `json:",omitempty"`
	// RDMATXBytes shall contain the total number of RDMA bytes transmitted on a
	// port since reset.
	//
	// Version added: v1.1.0
	RDMATXBytes *int `json:",omitempty"`
	// RDMATXReadRequests shall contain the total number of RDMA read requests
	// transmitted on a port since reset.
	//
	// Version added: v1.1.0
	RDMATXReadRequests *int `json:",omitempty"`
	// RDMATXRequests shall contain the total number of RDMA requests transmitted
	// on a port since reset.
	//
	// Version added: v1.1.0
	RDMATXRequests *int `json:",omitempty"`
	// RDMATXSendRequests shall contain the total number of RDMA send requests
	// transmitted on a port since reset.
	//
	// Version added: v1.1.0
	RDMATXSendRequests *int `json:",omitempty"`
	// RDMATXWriteRequests shall contain the total number of RDMA write requests
	// transmitted on a port since reset.
	//
	// Version added: v1.1.0
	RDMATXWriteRequests *int `json:",omitempty"`
	// RXBroadcastFrames shall contain the total number of valid broadcast frames
	// received on a port since reset, including host and remote management
	// passthrough traffic.
	//
	// Version added: v1.1.0
	RXBroadcastFrames *int `json:",omitempty"`
	// RXDiscards shall contain the total number of frames discarded in a port's
	// receive path since reset.
	//
	// Version added: v1.1.0
	RXDiscards *int `json:",omitempty"`
	// RXFCSErrors shall contain the total number of frames received with frame
	// check sequence (FCS) errors on a port since reset.
	//
	// Version added: v1.1.0
	RXFCSErrors *int `json:",omitempty"`
	// RXFalseCarrierErrors shall contain the total number of false carrier errors
	// received from phy on a port since reset.
	//
	// Version added: v1.1.0
	RXFalseCarrierErrors *int `json:",omitempty"`
	// RXFrameAlignmentErrors shall contain the total number of frames received
	// with alignment errors on a port since reset.
	//
	// Version added: v1.1.0
	RXFrameAlignmentErrors *int `json:",omitempty"`
	// RXFrames shall contain the total number of frames received on a port since
	// reset.
	//
	// Version added: v1.1.0
	RXFrames *int `json:",omitempty"`
	// RXMulticastFrames shall contain the total number of valid multicast frames
	// received on a port since reset, including host and remote management
	// passthrough traffic.
	//
	// Version added: v1.1.0
	RXMulticastFrames *int `json:",omitempty"`
	// RXOversizeFrames shall contain the total number of frames that exceed the
	// maximum frame size.
	//
	// Version added: v1.1.0
	RXOversizeFrames *int `json:",omitempty"`
	// RXPFCFrames shall contain the total number of priority flow control (PFC)
	// frames received on a port since reset.
	//
	// Version added: v1.1.0
	RXPFCFrames *int `json:",omitempty"`
	// RXPauseXOFFFrames shall contain the total number of flow control frames from
	// the network to pause transmission.
	//
	// Version added: v1.1.0
	RXPauseXOFFFrames *int `json:",omitempty"`
	// RXPauseXONFrames shall contain the total number of flow control frames from
	// the network to resume transmission.
	//
	// Version added: v1.1.0
	RXPauseXONFrames *int `json:",omitempty"`
	// RXUndersizeFrames shall contain the total number of frames that are smaller
	// than the minimum frame size of 64 bytes.
	//
	// Version added: v1.1.0
	RXUndersizeFrames *int `json:",omitempty"`
	// RXUnicastFrames shall contain the total number of valid unicast frames
	// received on a port since reset.
	//
	// Version added: v1.1.0
	RXUnicastFrames *int `json:",omitempty"`
	// TXBroadcastFrames shall contain the total number of good broadcast frames
	// transmitted on a port since reset, including host and remote management
	// passthrough traffic.
	//
	// Version added: v1.1.0
	TXBroadcastFrames *int `json:",omitempty"`
	// TXDiscards shall contain the total number of frames discarded in a port's
	// transmit path since reset.
	//
	// Version added: v1.1.0
	TXDiscards *int `json:",omitempty"`
	// TXExcessiveCollisions shall contain the number of times a single transmitted
	// frame encountered more than 15 collisions.
	//
	// Version added: v1.1.0
	TXExcessiveCollisions *int `json:",omitempty"`
	// TXFrames shall contain the total number of frames transmitted on a port
	// since reset.
	//
	// Version added: v1.1.0
	TXFrames *int `json:",omitempty"`
	// TXLateCollisions shall contain the total number of collisions that occurred
	// after one slot time as defined by IEEE 802.3.
	//
	// Version added: v1.1.0
	TXLateCollisions *int `json:",omitempty"`
	// TXMulticastFrames shall contain the total number of good multicast frames
	// transmitted on a port since reset, including host and remote management
	// passthrough traffic.
	//
	// Version added: v1.1.0
	TXMulticastFrames *int `json:",omitempty"`
	// TXMultipleCollisions shall contain the times that a transmitted frame
	// encountered 2-15 collisions.
	//
	// Version added: v1.1.0
	TXMultipleCollisions *int `json:",omitempty"`
	// TXPFCFrames shall contain the total number of priority flow control (PFC)
	// frames sent on a port since reset.
	//
	// Version added: v1.1.0
	TXPFCFrames *int `json:",omitempty"`
	// TXPauseXOFFFrames shall contain the total number of XOFF frames transmitted
	// to the network.
	//
	// Version added: v1.1.0
	TXPauseXOFFFrames *int `json:",omitempty"`
	// TXPauseXONFrames shall contain the total number of XON frames transmitted to
	// the network.
	//
	// Version added: v1.1.0
	TXPauseXONFrames *int `json:",omitempty"`
	// TXSingleCollisions shall contain the times that a successfully transmitted
	// frame encountered a single collision.
	//
	// Version added: v1.1.0
	TXSingleCollisions *int `json:",omitempty"`
	// TXUnicastFrames shall contain the total number of good unicast frames
	// transmitted on a port since reset, including host and remote management
	// passthrough traffic.
	//
	// Version added: v1.1.0
	TXUnicastFrames *int `json:",omitempty"`
}

// SAS shall describe physical (phy) related metrics for Serial Attached SCSI
// (SAS).
type SAS struct {
	// InvalidDwordCount shall contain the number of invalid dwords that have been
	// received by the phy outside of phy reset sequences.
	//
	// Version added: v1.1.0
	InvalidDwordCount *int `json:",omitempty"`
	// LossOfDwordSynchronizationCount shall contain the number of times the phy
	// has restarted the link reset sequence because it lost dword synchronization.
	//
	// Version added: v1.1.0
	LossOfDwordSynchronizationCount *int `json:",omitempty"`
	// PhyResetProblemCount shall contain the number of times a phy reset problem
	// has occurred.
	//
	// Version added: v1.5.0
	PhyResetProblemCount *int `json:",omitempty"`
	// RunningDisparityErrorCount shall contain the number of dwords containing
	// running disparity errors that have been received by the phy outside of phy
	// reset sequences.
	//
	// Version added: v1.1.0
	RunningDisparityErrorCount *int `json:",omitempty"`
}

// Transceiver shall describe the transceiver-related metrics.
type Transceiver struct {
	// ByLane shall contain an array of lane-related metrics for a transceiver in
	// this port. Each member in the array shall represent a single lane.
	//
	// Version added: v1.8.0
	ByLane []TransceiverLaneMetrics
	// RXInputPowerMilliWatts shall contain the RX input power value of the
	// transceiver, aggregated across all lanes of this port.
	//
	// Version added: v1.1.0
	RXInputPowerMilliWatts *float64 `json:",omitempty"`
	// SupplyVoltage shall contain the supply voltage for the transceiver.
	//
	// Version added: v1.1.0
	SupplyVoltage *float64 `json:",omitempty"`
	// TXBiasCurrentMilliAmps shall contain the TX bias current value of the
	// transceiver, aggregated across all lanes of this port.
	//
	// Version added: v1.1.0
	TXBiasCurrentMilliAmps *float64 `json:",omitempty"`
	// TXOutputPowerMilliWatts shall contain the TX output power value of the
	// transceiver, aggregated across all lanes of this port.
	//
	// Version added: v1.1.0
	TXOutputPowerMilliWatts *float64 `json:",omitempty"`
	// WavelengthNanometers shall contain the laser wavelength, in nanometers, for
	// the transceiver. This property shall not be present for non-optic
	// transceiver mediums.
	//
	// Version added: v1.7.0
	WavelengthNanometers string
}

// TransceiverLaneMetrics shall describe transceiver-related metrics on a
// per-lane basis.
type TransceiverLaneMetrics struct {
	// LaneID shall contain the transceiver lane for which metrics are being
	// supplied.
	//
	// Version added: v1.8.0
	LaneID *uint `json:"LaneId,omitempty"`
	// RXInputPowerMilliWatts shall contain the RX input power value of the
	// transceiver for the lane identified as 'LaneId'.
	//
	// Version added: v1.8.0
	RXInputPowerMilliWatts *float64 `json:",omitempty"`
	// TXBiasCurrentMilliAmps shall contain the TX bias current value of the
	// transceiver for the lane identified as 'LaneId'.
	//
	// Version added: v1.8.0
	TXBiasCurrentMilliAmps *float64 `json:",omitempty"`
	// TXOutputPowerMilliWatts shall contain the TX output power value of the
	// transceiver for the lane identified as 'LaneId'.
	//
	// Version added: v1.8.0
	TXOutputPowerMilliWatts *float64 `json:",omitempty"`
}
