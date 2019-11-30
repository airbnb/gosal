package xpreports

import (
	"strconv"

	"github.com/airbnb/gosal/config"
	"github.com/airbnb/gosal/xpreports/linux"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
)

// buildReport creates a report using linux APIs and paths.
func buildMachineReport(conf *config.Config) (*Machine, error) {
	u1 := uuid.NewV4().String()

	host, err := host.Info()
	if err != nil {
		return nil, errors.Wrap(err, "reports: getting host information")
	}

	disk, err := linux.GetDisk()
	if err != nil {
		return nil, errors.Wrap(err, "reports: getting root volume")
	}

	serial, err := linux.GetlinuxSerial()
	if err != nil {
		return nil, errors.Wrap(err, "reports: getting serial")

	}

	cpu, err := linux.GetProcessor()
	if err != nil {
		return nil, errors.Wrap(err, "machineinfo/gethardware: failed getting processor data")
	}

	consoleUser, _ := linux.GetLoggedInUsers()

	v, _ := mem.VirtualMemory()
	h, _ := host.Info()

	computerSystem, _ := linux.GetLinuxComputerSystem()

	// Convert memory from kb to correct size
	convertedMemory := float64(v.Total)
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
		ExtraData: &MachineExtraData{
			SerialNumber:         serial,
			HostName:             host.Hostname,
			ConsoleUser:          consoleUser[0],
			OSFamily:             "Linux",
			OperatingSystem:      h.PlatformVersion,
			HDSpace:              disk.FreeSpace,
			HDTotal:              disk.Size,
			MachineModel:         computerSystem.Model,
			MachineModelFriendly: "N/A",
			CPUType:              cpu.CPUType,
			CPUSpeed:             cpu.CurrentProcessorSpeed,
			Memory:               strMemory,
			MemoryKB:             int(v.Total / 1024),
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
