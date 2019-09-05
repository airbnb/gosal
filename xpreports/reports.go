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
	Facts     *machineFacts     `json:"facts"`
	ExtraData *machineExtraData `json:"extra_data"`
}

// Sal blah
type Sal struct {
	ExtraData *salExtraData `json:"extra_data"`
	Facts     *salFacts     `json:"facts"`
}

type salExtraData struct {
	Key        string `json:"key"`
	SalVersion string `json:"sal_version"`
}

type salFacts struct {
	CheckinModuleVersion string `json:"checkin_module_version"`
}

type machineFacts struct {
	CheckinModuleVersion string `json:"checkin_module_version"`
}

type machineExtraData struct {
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
	CPUSpeed             string  `json:"cpu_speed"`
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
	salReport, err := buildSalReport(conf)

	report := &Report{
		Machine: machineReport,
		Sal:     salReport,
	}

	return report, err
}
