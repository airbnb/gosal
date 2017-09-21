package reports

import (
	"log"
)

func BuildBase64bz2Report() (Base64bz2Report, error) {

	cDrive, err := GetCDrive()
	if err != nil {
		// TODO return the error here?
		log.Printf("reports: getting win32 disk: %s", err)
	}

	report := Base64bz2Report{
			AvailableDiskSpace: cDrive.FreeSpace,
	}

	return report, nil
}

type Base64bz2Report struct {
	AvailableDiskSpace int
}

// this appears to be what sal is expecting as a top level item
// https://github.com/salopensource/sal/blob/master/server/views.py#L1926
// type Base64bz2Report struct {
// 	base64bz2report BaseReport
// }
