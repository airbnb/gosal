package windows

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

	// Convert memory from kb to correct size
	convertedMemory := float64(os.TotalVisibleMemorySize)
	unitCount := 0
	strMemory := ""

	for convertedMemory >= 1024 {
		convertedMemory = convertedMemory / 1024
		unitCount++
	}

	switch unitCount {
	case 0:
		strMemory = strconv.FormatFloat(convertedMemory, 'f', 0, 64) + " KB"
	case 1:
		strMemory = strconv.FormatFloat(convertedMemory, 'f', 0, 64) + " MB"
	case 2:
		strMemory = strconv.FormatFloat(convertedMemory, 'f', 0, 64) + " GB"
	case 3:
		strMemory = strconv.FormatFloat(convertedMemory, 'f', 0, 64) + " TB"
	}

	hwinfo := HardwareInfo{
		MachineModel:          computerSystem.Model,
		CPUType:               cpu.CPUType,
		CurrentProcessorSpeed: cpu.CurrentProcessorSpeed,
		PhysicalMemory:        strMemory,
	}

	return &hwinfo, nil
}
