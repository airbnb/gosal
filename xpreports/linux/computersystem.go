package linux

import (
	"github.com/airbnb/gosal/xpreports/common"
	"github.com/dselans/dmidecode"
)

// GetLinuxComputerSystem
func GetLinuxComputerSystem() (LinuxComputerSystem, error) {
	dmi := dmidecode.New()
	


if err := dmi.Run(); err != nil {
   
}

	byNameData, _ := dmi.SearchByName("Base Board Information")




	usernames, _ := common.GetLoggedInUsers()

	CompSys := LinuxComputerSystem{
		UserName:     usernames[0],
		Manufacturer: byNameData[0]["Manufacturer"],
		Model:        byNameData[0]["Product Name"],
	}

	return CompSys, nil
}

// Win32ComputerSystem structure
type LinuxComputerSystem struct {
	UserName     string
	Manufacturer string
	Model        string
}
