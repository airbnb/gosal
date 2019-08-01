package linux

import (
	"fmt"

	"github.com/dselans/dmidecode"
)

// GetWin32ComputerSystem exports win32_ComputerSystem powershell class
func GetWin32ComputerSystem() (LinuxComputerSystem, error) {
	var CompSys LinuxComputerSystem

	dmi := dmidecode.New()
	if err := dmi.Run(); err != nil {
		fmt.Printf("Unable to get dmidecode information. Error: %v\n", err)
	}

	// manufacturer, _ := dmi.SearchByName("System Manufacturer")
	// model, _ := dmi.SearchByName("System Version")

	//	CompSys.Model = string(model)
	//	CompSys.Manufacturer = string(manufacturer)

	return CompSys, nil
}

// Win32ComputerSystem structure
type LinuxComputerSystem struct {
	Manufacturer string
	Model        string
}
