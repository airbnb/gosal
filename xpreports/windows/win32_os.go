package windows

import (
	"encoding/json"
	"os/exec"

	"github.com/pkg/errors"
)

// GetWin32OS exports win32_operatingsystem powershell class
func GetWin32OS() (Win32OS, error) {
	cmd := exec.Command("powershell", "gwmi", "Win32_OperatingSystem", "|", "ConvertTo-Json")

	// cmd.Stderr = os.Stderr
	o, err := cmd.Output()
	if err != nil {
		return Win32OS{}, errors.Wrap(err, "exec gwmi Win32_OperatingSystem")
	}

	var j Win32OS

	if err := json.Unmarshal(o, &j); err != nil {
		return Win32OS{}, errors.Wrap(err, "failed unmarshalling Win32_OperatingSystem")
	}

	return j, nil
}

// Win32OS structure
type Win32OS struct {
	Caption                string `json:"Caption"` // os version
	TotalVirtualMemorySize int    `json:"TotalVirtualMemorySize"`
	TotalVisibleMemorySize int    `json:"TotalVisibleMemorySize"`
}

/*
[
    "",
    "",
    "SystemDirectory : C:\\WINDOWS\\system32",
    "Organization    : ",
    "BuildNumber     : 17763",
    "RegisteredUser  : Windows User",
    "SerialNumber    : 00331-20092-26760-AA304",
    "Version         : 10.0.17763",
    "",
    "",
    ""
]
*/
