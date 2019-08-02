package darwin

import (
	"fmt"

	"github.com/dselans/dmidecode"
)

// GetWMacOSComputerSystem exports  powershell class
func GetWMacOSComputerSystem() (MacOSComputerSystem, error) {
	var CompSys MacOSComputerSystem

	dmi := dmidecode.New()
	if err := dmi.Run(); err != nil {
		fmt.Printf("Unable to get dmidecode information. Error: %v\n", err)
	}

	return CompSys, nil
}

// Win32ComputerSystem structure
type MacOSComputerSystem struct {
	UserName     string
	Manufacturer string
	Model        string
}
