package reports

import (
	"strconv"

	"github.com/airbnb/gosal/config"
	"github.com/pkg/errors"
	"github.com/satori/go.uuid"
)

// BuildReport builds the report object
func BuildReport(conf *config.Config) (*Report, error) {

	win32Bios, err := GetWin32Bios()
	if err != nil {
		return nil, errors.Wrap(err, "get win32Bios")
	}

	CDrive, err := GetCDrive()
	if err != nil {
		return nil, errors.Wrap(err, "reports: getting win32 disk")
	}

	u1 := uuid.NewV4().String()

	encodedCompressedPlist, err := BuildBase64bz2Report(conf)
	if err != nil {
		return nil, errors.Wrap(err, "reports: getting plist")
	}

	report := &Report{
		Serial:          win32Bios.SerialNumber,
		Key:             conf.Key,
		Name:            win32Bios.PSComputerName,
		DiskSize:        strconv.Itoa(CDrive.Size),
		SalVersion:      strconv.Itoa(1),
		RunUUID:         u1,
		Base64bz2Report: encodedCompressedPlist,
	}

	// fmt.Printf("%+v\n", report)
	return report, nil
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
