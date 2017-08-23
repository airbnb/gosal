package reports

import (
	"encoding/json"
	"os/exec"
)

// GetWin32LogicalDisk returns an array of powershell class win32_logicaldisk
func GetWin32LogicalDisk() ([]Win32LogicalDisk, error) {
	cmd := exec.Command("powershell", "Get-WmiObject", "Win32_LogicalDisk", "|", "convertto-json")

	o, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	var j []Win32LogicalDisk

	if err := json.Unmarshal(o, &j); err != nil {
		return nil, err
	}

	return j, nil
}

// Win32LogicalDisk structure
type Win32LogicalDisk struct {
	Name      string `json:"Name"`
	Size      int    `json:"Size"`
	FreeSpace int64  `json:"Free"`
}
