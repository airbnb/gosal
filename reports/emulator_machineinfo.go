package reports

import (
	"log"
	"strconv"
)

// EmulateMachineInfo copies its behavior from macOS, and provides struct data to Sal
func EmulateMachineInfo() (MachineInfo, error) {

	win32OS, err := GetWin32OS()
	if err != nil {
		// TODO return the error here?
		log.Printf("reports: getting win32 os: %s", err)
	}

	hardwareInfo, err := GetHardwareInfo()
	if err != nil {
		// TODO return the error here?
		log.Printf("reports: system profile failed: %s", err)
	}

	report := MachineInfo{
		OSVers:       win32OS.Caption,
		HardwareInfo: hardwareInfo,
	}

	return report, nil
}

// MachineInfo is required as a top level report field
type MachineInfo struct {
	OSVers       string
	HardwareInfo HardwareInfo
}

// HardwareInfo is a subset of MachineInfo
type HardwareInfo struct {
	MachineModel          string
	CPUType               string
	CurrentProcessorSpeed int
	PhysicalMemory        string
}

// GetHardwareInfo creates the necessary structure sal expects
func GetHardwareInfo() (HardwareInfo, error) {

	computerSystem, err := GetWin32ComputerSystem()
	if err != nil {
		// TODO return the error here?
		log.Printf("machine info: computer system information failed: %s", err)
	}

	os, err := GetWin32OS()
	if err != nil {
		// TODO return the error here?
		log.Printf("machine info: os information failed: %s", err)
	}

	cpu, err := GetWin32Processor()
	if err != nil {
		// TODO return the error here?
		log.Printf("machine info: getting processor information failed: %s", err)
	}

	hwinfo := HardwareInfo{
		MachineModel:          computerSystem.Model,
		CPUType:               cpu.CPUType,
		CurrentProcessorSpeed: cpu.CurrentProcessorSpeed,
		PhysicalMemory:        strconv.Itoa(os.TotalVisibleMemorySize) + " KB",
	}

	return hwinfo, nil
}
