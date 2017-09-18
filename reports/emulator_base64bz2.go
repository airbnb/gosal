package reports

import (
	"log"
)

func BuildBase64bz2Report() (Base64bz2Report, error) {

	machineInfo, err := EmulateMachineInfo()
	if err != nil {
		// TODO return the error here?
		log.Printf("reports: machine info: %s", err)
	}

	cDrive, err := GetCDrive()
	if err != nil {
		// TODO return the error here?
		log.Printf("reports: getting win32 disk: %s", err)
	}

	report := Base64bz2Report{
			AvailableDiskSpace: cDrive.FreeSpace,
			MachineInfo:        machineInfo,
	}

	return report, nil
}

type Base64bz2Report struct {
	AvailableDiskSpace int
	MachineInfo        MachineInfo
}

// this appears to be what sal is expecting as a top level item
// https://github.com/salopensource/sal/blob/master/server/views.py#L1926
// type Base64bz2Report struct {
// 	base64bz2report BaseReport
// }
