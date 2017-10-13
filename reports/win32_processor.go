package reports

import (
	"encoding/json"
	"os/exec"
)

// GetWin32Processor exports the win32_bios powershell class
func GetWin32Processor() (Win32Processor, error) {
	cmd := exec.Command("powershell", "gwmi", "win32_processor", "|", "convertto-json")

	o, err := cmd.Output()
	if err != nil {
		return Win32Processor{}, err
	}

	var j Win32Processor

	if err := json.Unmarshal(o, &j); err != nil {
		return Win32Processor{}, err
	}

	return j, nil
}

// Win32Processor data structure
type Win32Processor struct {
	CPUType               string `json:"Name"`
	CurrentProcessorSpeed int    `json:"MaxClockSpeed"`
}
