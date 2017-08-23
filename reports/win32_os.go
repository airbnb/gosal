package reports

import (
	"encoding/json"
	"fmt"
	"os/exec"
)

// GetWin32OS exports win32_operatingsystem powershell class
func GetWin32OS() (Win32OS, error) {
	cmd := exec.Command("powershell", "gwmi", "Win32_OperatingSystem", "|", "convertto-json")

	o, err := cmd.Output()

	if err != nil {
		fmt.Println("error here")
		return Win32OS{}, err
	}

	var j Win32OS

	fmt.Println(j)

	if err := json.Unmarshal(o, &j); err != nil {
		return Win32OS{}, err
	}

	return j, nil
}

// Win32OS structure
type Win32OS struct {
	Caption                string `json:"Caption"` //os version
	TotalVirtualMemorySize int    `json:"TotalVirtualMemorySize"`
	TotalVisibleMemorySize int    `json:"TotalVisibleMemorySize"`
}
