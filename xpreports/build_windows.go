package xpreports

import (
	"strconv"

	"github.com/airbnb/gosal/config"
	"github.com/airbnb/gosal/xpreports/windows"
	"github.com/pkg/errors"
)

// buildReport creates the necessary struct for Machine
func buildMachineReport(conf *config.Config) (*Machine, error) {
	bios, err := windows.GetWin32Bios()
	if err != nil {
		return nil, errors.Wrap(err, "machineinfo/gethardware: failed getting bios data")
	}

	computerSystem, err := windows.GetWin32ComputerSystem()
	if err != nil {
		return nil, errors.Wrap(err, "machineinfo/gethardware: failed getting system data")
	}

	os, err := windows.GetWin32OS()
	if err != nil {
		return nil, errors.Wrap(err, "machineinfo/gethardware: failed getting os data")
	}

	cpu, err := windows.GetWin32Processor()
	if err != nil {
		return nil, errors.Wrap(err, "machineinfo/gethardware: failed getting processor data")
	}

	disk, err := windows.GetCDrive()
	if err != nil {
		return nil, errors.Wrap(err, "machineinfo/gethardware: failed getting information for c drive")
	}

	// Convert memory from kb to correct size
	convertedMemory := float64(os.TotalVisibleMemorySize)
	var unitCount int
	var strMemory string

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

	m := &Machine{
		ExtraData: &MachineExtraData{
			SerialNumber:         bios.SerialNumber,
			HostName:             bios.PSComputerName,
			ConsoleUser:          computerSystem.UserName,
			OSFamily:             "Windows",
			OperatingSystem:      os.Caption + " " + os.Version,
			HDSpace:              disk.FreeSpace,
			HDTotal:              disk.Size,
			MachineModel:         computerSystem.Model,
			MachineModelFriendly: "N/A",
			CPUType:              cpu.CPUType,
			CPUSpeed:             cpu.CPUSpeed,
			Memory:               strMemory,
			MemoryKB:             os.TotalVisibleMemorySize,
		}, Facts: &MachineFacts{
			CheckinModuleVersion: "1",
		},
	}

	return m, nil
}

func buildSalReport(conf *config.Config) (*Sal, error) {
	s := &Sal{
		ExtraData: &SalExtraData{
			Key:        conf.Key,
			SalVersion: "1",
		}, Facts: &SalFacts{
			CheckinModuleVersion: "1",
		},
	}

	return s, nil
}
