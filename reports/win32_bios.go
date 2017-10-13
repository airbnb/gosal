package reports

import (
	"encoding/json"
	"os/exec"

	"github.com/pkg/errors"
)

// GetWin32Bios exports the win32_bios powershell class
func GetWin32Bios() (Win32Bios, error) {
	cmd := exec.Command("powershell", "gwmi", "Win32_Bios", "|", "ConvertTo-Json")

	// cmd.Stderr = os.Stderr
	o, err := cmd.Output()
	if err != nil {
		return Win32Bios{}, errors.Wrap(err, "exec gwmi Win32_Bios")
	}

	var j Win32Bios

	if err := json.Unmarshal(o, &j); err != nil {
		return Win32Bios{}, errors.Wrap(err, "failed unmarshalling Win32_Bios")
	}

	return j, nil
}

// Win32Bios data structure
type Win32Bios struct {
	PSComputerName string `json:"PSComputerName"`
	SerialNumber   string `json:"SerialNumber"`
}
