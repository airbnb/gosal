package reports

import (
	"strconv"

	"github.com/pkg/errors"
)

// EmulateMachineInfo copies its behavior from macOS, and provides struct data to Sal
func EmulateMachineInfo() (*MachineInfo, error) {

	win32OS, err := GetWin32OS()
	if err != nil {
		return nil, errors.Wrap(err, "emulatemachineinfo: failed getting os data")
	}

	hardwareInfo, err := GetHardwareInfo()
	if err != nil {
		return nil, errors.Wrap(err, "emulatemachineinfo: failed getting hardware data")
	}

	report := MachineInfo{
		OSVers:       win32OS.Caption,
		HardwareInfo: hardwareInfo,
	}

	return &report, nil
}

// MachineInfo is required as a top level report field
type MachineInfo struct {
	OSVers       string
	HardwareInfo *HardwareInfo
}

// HardwareInfo is a subset of MachineInfo
type HardwareInfo struct {
	MachineModel          string
	CPUType               string
	CurrentProcessorSpeed int
	PhysicalMemory        string
}

// GetHardwareInfo creates the necessary structure sal expects
func GetHardwareInfo() (*HardwareInfo, error) {

	computerSystem, err := GetWin32ComputerSystem()
	if err != nil {
		return nil, errors.Wrap(err, "machineinfo/gethardware: failed getting system data")
	}

	os, err := GetWin32OS()
	if err != nil {
		return nil, errors.Wrap(err, "machineinfo/gethardware: failed getting os data")
	}

	cpu, err := GetWin32Processor()
	if err != nil {
		return nil, errors.Wrap(err, "machineinfo/gethardware: failed getting processor data")
	}

	hwinfo := HardwareInfo{
		MachineModel:          computerSystem.Model,
		CPUType:               cpu.CPUType,
		CurrentProcessorSpeed: cpu.CurrentProcessorSpeed,
		PhysicalMemory:        strconv.Itoa(os.TotalVisibleMemorySize) + " KB",
	}

	return &hwinfo, nil
}
