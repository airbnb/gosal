package xpreports

import (
	"github.com/airbnb/gosal/config"
)

// Report is a high level struct with individual checkin segments
type Report struct {
	Machine *Machine
	Sal     *Sal
}

// Machine blah
type Machine struct {
	Facts     *MachineFacts     `json:"facts"`
	ExtraData *MachineExtraData `json:"extra_data"`
}

// Sal blah
type Sal struct {
	ExtraData *SalExtraData `json:"extra_data"`
	Facts     *SalFacts     `json:"facts"`
}

// SalExtraData blah
type SalExtraData struct {
	Key        string `json:"key"`
	SalVersion string `json:"sal_version"`
}

// SalFacts blah
type SalFacts struct {
	CheckinModuleVersion string `json:"checkin_module_version"`
}

// MachineFacts blah
type MachineFacts struct {
	CheckinModuleVersion string `json:"checkin_module_version"`
}

// MachineExtraData blah
type MachineExtraData struct {
	SerialNumber         string  `json:"serial"`
	HostName             string  `json:"hostname"`
	ConsoleUser          string  `json:"console_user"`
	OSFamily             string  `json:"os_family"`
	OperatingSystem      string  `json:"operating_system"`
	HDSpace              int     `json:"hd_space"`
	HDTotal              int     `json:"hd_total"`
	HDPercent            float32 `json:"hd_percent"`
	MachineModel         string  `json:"machine_model"`
	MachineModelFriendly string  `json:"machine_model_friendly"`
	CPUType              string  `json:"cpu_type"`
	CPUSpeed             int     `json:"cpu_speed"`
	Memory               string  `json:"memory"`
	MemoryKB             int     `json:"memory_kb"`
}

// Build creates a report for the sal server.
// Build supports darwin, windows and linux and will use
// the appropriate APIs for each system.
func Build(conf *config.Config) (*Report, error) {
	// buildReport is implented separately for each
	// operating system.
	machineReport, err := buildMachineReport(conf)
	if err != nil {
		return nil, err
	}
	salReport, err := buildSalReport(conf)
	if err != nil {
		return nil, err
	}
	report := &Report{
		Machine: machineReport,
		Sal:     salReport,
	}

	return report, err
}
