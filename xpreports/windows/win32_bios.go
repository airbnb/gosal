package windows

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

/*
owershell gwmi Win32_Bios |ConvertTo-Json
[
    "",
    "",
    "SMBIOSBIOSVersion : VMW71.00V.12343141.B64.1902160724",
    "Manufacturer      : VMware, Inc.",
    "Name              : VMW71.00V.12343141.B64.1902160724",
    "SerialNumber      : VMware-56 4d ea bd d1 c3 fe 6d-80 7a ba 09 22 3f ba 82",
    "Version           : INTEL  - 6040000",
    "",
    "",
    ""
]
*/
