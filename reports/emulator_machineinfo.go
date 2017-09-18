package reports

import (
	"log"
)

func EmulateMachineInfo() (MachineInfo, error) {

	win32OS, err := GetWin32OS()
	if err != nil {
		// TODO return the error here?
		log.Printf("reports: getting win32 os: %s", err)
	}

	report := MachineInfo{
		os_vers:    win32OS.Caption,
	}

	return report, nil
}

type MachineInfo struct {
	os_vers    string
}
