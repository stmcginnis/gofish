//
// SPDX-License-Identifier: BSD-3-Clause
//

package schemas

import "encoding/json"

// UnmarshalJSON unmarshals a Processor object from the raw JSON.
func (p *Processor) UnmarshalJSON(b []byte) error {
	type temp Processor
	type pActions struct {
		Reset           ActionTarget `json:"#Processor.Reset"`
		ResetToDefaults ActionTarget `json:"#Processor.ResetToDefaults"`
	}
	type pLinks struct {
		Chassis                Link  `json:"Chassis"`
		ConnectedProcessors    Links `json:"ConnectedProcessors"`
		Endpoints              Links `json:"Endpoints"`
		FabricAdapters         Links `json:"FabricAdapters"`
		GraphicsController     Link  `json:"GraphicsController"`
		Memory                 Links `json:"Memory"`
		NetworkDeviceFunctions Links `json:"NetworkDeviceFunctions"`
		PCIeDevice             Link  `json:"PCIeDevice"`
		PCIeFunctions          Links `json:"PCIeFunctions"`
	}
	var tmp struct {
		temp
		Actions                pActions
		Links                  pLinks
		AccelerationFunctions  Link                             `json:"AccelerationFunctions"`
		AppliedOperatingConfig ResolvedLink[OperatingConfig]    `json:"AppliedOperatingConfig"`
		Assembly               ResolvedLink[Assembly]           `json:"Assembly"`
		CacheMemory            Link                             `json:"CacheMemory"`
		Certificates           Link                             `json:"Certificates"`
		EnvironmentMetrics     ResolvedLink[EnvironmentMetrics] `json:"EnvironmentMetrics"`
		Metrics                Link                             `json:"Metrics"`
		OperatingConfigs       Link                             `json:"OperatingConfigs"`
		Ports                  Link                             `json:"Ports"`
		SubProcessors          Link                             `json:"SubProcessors"`

		MaxSpeedMHz any
		Socket      any // ZTE bug
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*p = Processor(tmp.temp)

	// Extract the links to other entities for later
	p.resetTarget = tmp.Actions.Reset.Target
	p.resetToDefaultsTarget = tmp.Actions.ResetToDefaults.Target
	p.chassis = tmp.Links.Chassis.String()
	p.connectedProcessors = tmp.Links.ConnectedProcessors.ToStrings()
	p.endpoints = tmp.Links.Endpoints.ToStrings()
	p.fabricAdapters = tmp.Links.FabricAdapters.ToStrings()
	p.graphicsController = tmp.Links.GraphicsController.String()
	p.memory = tmp.Links.Memory.ToStrings()
	p.networkDeviceFunctions = tmp.Links.NetworkDeviceFunctions.ToStrings()
	p.pCIeDevice = tmp.Links.PCIeDevice.String()
	p.pCIeFunctions = tmp.Links.PCIeFunctions.ToStrings()
	p.accelerationFunctions = tmp.AccelerationFunctions.String()
	p.appliedOperatingConfig = tmp.AppliedOperatingConfig

	p.assembly = tmp.Assembly

	p.cacheMemory = tmp.CacheMemory.String()
	p.certificates = tmp.Certificates.String()
	p.environmentMetrics = tmp.EnvironmentMetrics

	p.metrics = tmp.Metrics.String()
	p.operatingConfigs = tmp.OperatingConfigs.String()
	p.ports = tmp.Ports.String()
	p.subProcessors = tmp.SubProcessors.String()

	p.MaxSpeedMHz = toInt(tmp.MaxSpeedMHz)
	p.Socket = parseString(tmp.Socket)

	// This is a read/write object, so we need to save the raw object data for later
	p.RawData = b

	return nil
}

// Metrics gets the Metrics linked resource.
func (p *Processor) Metrics() (*ProcessorMetrics, error) {
	if p.metrics == "" {
		return nil, nil
	}
	return GetObject[ProcessorMetrics](p.client, p.metrics)
}

// UnmarshalJSON unmarshals a MemorySummary object from the raw JSON.
func (m *ProcessorMemorySummary) UnmarshalJSON(b []byte) error {
	type temp ProcessorMemorySummary
	var tmp struct {
		temp
		Metrics Link `json:"Metrics"`
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*m = ProcessorMemorySummary(tmp.temp)

	// Extract the links to other entities for later
	m.metrics = tmp.Metrics.String()

	return nil
}

// Metrics gets the Metrics linked resource.
func (m *ProcessorMemorySummary) Metrics(client Client) (*MemoryMetrics, error) {
	if m.metrics == "" {
		return nil, nil
	}
	return GetObject[MemoryMetrics](client, m.metrics)
}
