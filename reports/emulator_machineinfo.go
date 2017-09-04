package reports

import (
	"log"
)

func EmulateMachineInfo() (MachineInfo, error) {

	win32Bios, err := GetWin32Bios()
	if err != nil {
		// TODO return the error here?
		log.Printf("reports: getting win32 bios: %s", err)
	}

	win32ComputerSystem, err := GetWin32ComputerSystem()
	if err != nil {
		// TODO return the error here?
		log.Printf("reports: getting win32 computer system: %s", err)
	}

	win32OS, err := GetWin32OS()
	if err != nil {
		// TODO return the error here?
		log.Printf("reports: getting win32 os: %s", err)
	}

	report := MachineInfo{
		Hostname:     win32Bios.PSComputerName,
		MachineModel: win32ComputerSystem.Model,
		OSVersion:    win32OS.Caption,
		SerialNumber: win32Bios.SerialNumber,
	}

	return report, nil

}

type MachineInfo struct {
	Hostname     string
	MachineModel string
	OSVersion    string
	SerialNumber string
}
