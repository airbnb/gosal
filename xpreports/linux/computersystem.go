package linux

import (
	"github.com/dselans/dmidecode"
)

// GetComputerSystem
func GetComputerSystem() (LinuxComputerSystem, error) {
	dmi := dmidecode.New()

	_ = dmi.Run()

	byNameData, err := dmi.SearchByName("System Information")
	if err != nil {
		return LinuxComputerSystem{}, err
	}

	usernames, err := ConsoleUser()
	if err != nil {
		return LinuxComputerSystem{}, err
	}

	CompSys := LinuxComputerSystem{
		UserName:     usernames[0],
		Manufacturer: byNameData[0]["Manufacturer"],
		Model:        byNameData[0]["Product Name"],
	}

	return CompSys, nil
}

// LinucComputerSystem
type LinuxComputerSystem struct {
	UserName     string
	Manufacturer string
	Model        string
}
