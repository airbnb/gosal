package xpreports

import (
	"strconv"

	"github.com/airbnb/gosal/config"
	"github.com/bdemetris/gosal/xpreports/windows"
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

	m := &Machine{
		ExtraData: &machineExtraData{
			SerialNumber:    bios.SerialNumber,
			MachineModel:    computerSystem.Model,
			CPUType:         cpu.CPUType,
			Memory:          strMemory,
			OperatingSystem: os.Caption,
		}, Facts: &machineFacts{
			CheckinModuleVersion: "1",
		},
	}

	return m, nil
}

func buildSalReport(conf *config.Config) (*Sal, error) {
	s := &Sal{
		ExtraData: &salExtraData{
			Key:        conf.Key,
			SalVersion: "1",
		}, Facts: &salFacts{
			CheckinModuleVersion: "1",
		},
	}

	return s, nil
}
