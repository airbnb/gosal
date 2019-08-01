package darwin

import (
	"fmt"

	"github.com/dselans/dmidecode"
)

// GetWin32ComputerSystem exports win32_ComputerSystem powershell class
func GetWin32ComputerSystem() (MacOSComputerSystem, error) {
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
