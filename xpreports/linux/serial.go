package linux

import (
	"github.com/dselans/dmidecode"
	"github.com/pkg/errors"
)

// GetlinuxSerial returns the system serial number
func GetlinuxSerial() (string, error) {
	dmi := dmidecode.New()

	err := dmi.Run()
	if err != nil {
		return "", errors.Wrap(err, "DMI run")
	}

	byNameData, err := dmi.SearchByName("System Information")
	if err != nil {
		return "", errors.Wrap(err, "extracting information from DMI data")
	}

	return byNameData[0]["Serial Number"], nil
}
