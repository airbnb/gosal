package reports

import (
	"encoding/json"
	"os/exec"
)

// GetWin32Bios exports the win32_bios powershell class
func GetWin32Bios() (Win32Bios, error) {
	cmd := exec.Command("powershell", "gwmi", "win32_bios", "|", "convertto-json")

	o, err := cmd.Output()
	if err != nil {
		return Win32Bios{}, err
	}

	var j Win32Bios

	if err := json.Unmarshal(o, &j); err != nil {
		return Win32Bios{}, err
	}

	return j, nil
}

// Win32Bios data structure
type Win32Bios struct {
	PSComputerName string `json:"PSComputerName"`
	SerialNumber   string `json:"SerialNumber"`
}
