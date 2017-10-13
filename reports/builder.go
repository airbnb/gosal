package reports

import (
	"log"
	"strconv"

	"github.com/satori/go.uuid"
)

// BuildReport builds the report object
func BuildReport(apiKey string) Report {

	win32Bios, _ := GetWin32Bios()
	CDrive, err := GetCDrive()
	if err != nil {
		// TODO return the error here?
		log.Printf("reports: getting win32 disk: %s", err)
	}

	u1 := uuid.NewV4().String()

	encodedCompressedPlist, err := BuildBase64bz2Report()
	if err != nil {
		// TODO return the error here?
		log.Printf("reports: getting plist: %s", err)
	}

	report := Report{
		Serial:          win32Bios.SerialNumber,
		Key:             apiKey,
		Name:            win32Bios.PSComputerName,
		DiskSize:        strconv.Itoa(CDrive.Size),
		SalVersion:      strconv.Itoa(1),
		RunUUID:         u1,
		Base64bz2Report: encodedCompressedPlist,
	}

	// fmt.Printf("%+v\n", report)
	return report
}

// Report structure
type Report struct {
	Serial          string
	Key             string
	Name            string
	DiskSize        string
	SalVersion      string
	RunUUID         string
	Base64bz2Report string
}
