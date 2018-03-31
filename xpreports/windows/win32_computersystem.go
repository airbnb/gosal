package windows

import (
	"encoding/json"
	"os/exec"

	"github.com/pkg/errors"
)

// GetWin32ComputerSystem exports win32_ComputerSystem powershell class
func GetWin32ComputerSystem() (Win32ComputerSystem, error) {
	cmd := exec.Command("powershell", "gwmi", "Win32_ComputerSystem", "|", "ConvertTo-Json")

	// cmd.Stderr = os.Stderr
	o, err := cmd.Output()
	if err != nil {
		return Win32ComputerSystem{}, errors.Wrap(err, "exec gwmi Win32_ComputerSystem")
	}

	var j Win32ComputerSystem

	if err := json.Unmarshal(o, &j); err != nil {
		return Win32ComputerSystem{}, errors.Wrap(err, "failed unmarshalling Win32_ComputerSystem")
	}

	return j, nil
}

// Win32ComputerSystem structure
type Win32ComputerSystem struct {
	UserName     string `json:"UserName"`
	Manufacturer string `json:"Manufacturer"`
	Model        string `json:"Model"`
}
