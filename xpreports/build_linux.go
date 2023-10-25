package xpreports

import (
	"strconv"

	"github.com/airbnb/gosal/config"
	"github.com/airbnb/gosal/xpreports/linux"
	"github.com/pkg/errors"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/mem"
)

// buildReport creates a report using linux APIs and paths.
func buildMachineReport(conf *config.Config) (*Machine, error) {
	host, err := host.Info()
	if err != nil {
		return nil, errors.Wrap(err, "reports: getting host information")
	}

	disk, err := linux.Disk()
	if err != nil {
		return nil, errors.Wrap(err, "reports: getting root volume")
	}

	serial, err := linux.Serial()
	if err != nil {
		return nil, errors.Wrap(err, "reports: getting serial")
	}

	cpu, err := linux.GetProcessor()
	if err != nil {
		return nil, errors.Wrap(err, "machineinfo/gethardware: failed getting processor data")
	}

	consoleUser, err := linux.ConsoleUser()
	if err != nil {
		return nil, errors.Wrap(err, "reports: getting console user")
	}

	//Return the underling linux meminfo detail in bytes
	v, _ := mem.VirtualMemory()

	computerSystem, err := linux.GetComputerSystem()
	if err != nil {
		return nil, errors.Wrap(err, "reports: getting computerSystem")
	}

	// Convert memory from bytes to a human readable size format
	convertedMemory := float64(v.Total)
	var unitCount int
	var strMemory string

	for convertedMemory >= 1024 {
		convertedMemory = convertedMemory / 1024
		unitCount++
	}

	switch unitCount {
	case 0:
		strMemory = strconv.FormatFloat(convertedMemory, 'f', 0, 64) + " B"
	case 1:
		strMemory = strconv.FormatFloat(convertedMemory, 'f', 0, 64) + " KB"
	case 2:
		strMemory = strconv.FormatFloat(convertedMemory, 'f', 0, 64) + " MB"
	case 3:
		strMemory = strconv.FormatFloat(convertedMemory, 'f', 0, 64) + " GB"
	case 4:
		strMemory = strconv.FormatFloat(convertedMemory, 'f', 0, 64) + " TB"
	}

	m := &Machine{
		ExtraData: &MachineExtraData{
			SerialNumber:         serial,
			HostName:             host.Hostname,
			ConsoleUser:          consoleUser[0],
			OSFamily:             "Linux",
			OperatingSystem:      host.PlatformVersion,
			HDSpace:              disk.FreeSpace,
			HDTotal:              disk.Size,
			HDPercent:            disk.PercentageFree,
			MachineModel:         computerSystem.Model,
			MachineModelFriendly: "N/A",
			CPUType:              cpu.CPUType,
			CPUSpeed:             cpu.CurrentProcessorSpeed,
			Memory:               strMemory,
			MemoryKB:             int(v.Total) / 1024,
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
