package linux

import (
	"github.com/dselans/dmidecode"
)

// GetlinuxSerial returns the system serial number
func GetlinuxSerial() (string, error) {
	dmi := dmidecode.New()

	err := dmi.Run()
	if err != nil {
		return "", err
	}

	byNameData, err := dmi.SearchByName("System Information")
	if err != nil {
		return "", err
	}

	return byNameData[0]["Serial Number"], nil
}
