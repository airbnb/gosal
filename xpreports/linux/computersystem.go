package linux

import (
	"github.com/dselans/dmidecode"
)

// GetLinuxComputerSystem
func GetLinuxComputerSystem() (LinuxComputerSystem, error) {
	dmi := dmidecode.New()

	_ = dmi.Run()

	byNameData, _ := dmi.SearchByName("System Information")

	usernames, _ := GetLoggedInUsers()

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
