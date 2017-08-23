package reports

import (
	"encoding/json"
	"os/exec"
)

// GetWin32ComputerSystem exports win32_ComputerSystem powershell class
func GetWin32ComputerSystem() (Win32ComputerSystem, error) {
	cmd := exec.Command("powershell", "gwmi", "win32_ComputerSystem", "|", "ConvertTo-Json")

	o, err := cmd.Output()
	if err != nil {
		return Win32ComputerSystem{}, err
	}

	var j Win32ComputerSystem

	if err := json.Unmarshal(o, &j); err != nil {
		return Win32ComputerSystem{}, err
	}

	return j, nil
}

// Win32ComputerSystem structure
type Win32ComputerSystem struct {
	UserName string `json:"UserName"`
}
