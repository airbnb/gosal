package windows

import (
	"encoding/json"
	"os/exec"

	"github.com/pkg/errors"
)

// GetWin32LogicalDisk returns an array of powershell class win32_logicaldisk
func GetWin32LogicalDisk() ([]Win32LogicalDisk, error) {
	cmd := exec.Command("powershell", "gwmi", "Win32_LogicalDisk", "|", "ConvertTo-Json")

	// cmd.Stderr = os.Stderr
	o, err := cmd.Output()
	if err != nil {
		return nil, errors.Wrap(err, "exec gwmi Win32_LogicalDisk")
	}

	var j []Win32LogicalDisk

	if err := json.Unmarshal(o, &j); err != nil {
		return nil, errors.Wrap(err, "failed unmarshalling Win32LogicalDisk")
	}

	return j, nil
}

// Win32LogicalDisk structure
type Win32LogicalDisk struct {
	Name      string `json:"Name"`
	Size      int    `json:"Size"`
	FreeSpace int    `json:"FreeSpace"`
}

// GetCDrive explicity looks for C Drive
func GetCDrive() (Win32LogicalDisk, error) {

	disks, _ := GetWin32LogicalDisk()

	var c Win32LogicalDisk

	for _, element := range disks {
		if element.Name == "C:" {
			c.Name = element.Name
			c.Size = (element.Size / 1024)
			c.FreeSpace = (element.FreeSpace / 1024)
		}
	}

	return c, nil
}
