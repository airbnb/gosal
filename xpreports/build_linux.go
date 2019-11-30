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

	disk, err := common.GetDisk()
	if err != nil {
		return nil, errors.Wrap(err, "reports: getting root volume")
	}

	serial, err := linux.GetlinuxSerial()
	if err != nil {
		return nil, errors.Wrap(err, "reports: getting serial")
	}

	encodedCompressedPlist, err := linux.BuildBase64bz2Report(conf)
	if err != nil {
		return nil, errors.Wrap(err, "reports: getting plist")
	}

	consoleUser, err := GetLoggedInUsers()
	if err != nil {
		return "", errors.Wrap(err, "Getting logged in user")
	}

	v, _ := mem.VirtualMemory()
	h, _ := host.Info()

	computerSystem, err := GetLinuxComputerSystem()
	if err != nil {
		return nil, errors.Wrap(err, "machineinfo/gethardware: failed getting system data")
	}

	// Convert memory from kb to correct size
	convertedMemory := float64(memory.TotalVisibleMemorySize)
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
			SerialNumber:         serial,                    // done
			HostName:             host.Hostname,             // done
			ConsoleUser:          consoleUser[0],            // done
			OSFamily:             "Linux",                   // done
			OperatingSystem:      h.PlatformVersion,         // done
			HDSpace:              disk.FreeSpace,            // done
			HDTotal:              disk.Size,                 // done
			MachineModel:         computerSystem.Model,      // done
			MachineModelFriendly: "N/A",                     // done
			CPUType:              cpu.CPUType,               // done
			CPUSpeed:             cpu.CurrentProcessorSpeed, // done
			Memory:               strMemory,                 // done
			MemoryKB:             v.Total / 1024,            // done
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
