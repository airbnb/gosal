package windows

import (
	"encoding/json"
	"os/exec"

	"github.com/pkg/errors"
)

// GetWin32Processor exports the win32_bios powershell class
func GetWin32Processor() (Win32Processor, error) {
	cmd := exec.Command("powershell", "gwmi", "Win32_Processor", "|", "ConvertTo-Json")

	// cmd.Stderr = os.Stderr
	o, err := cmd.Output()
	if err != nil {
		return Win32Processor{}, errors.Wrap(err, "gwmi exec Win32_Processor")
	}

	var j Win32Processor

	if err := json.Unmarshal(o, &j); err != nil {
		return Win32Processor{}, errors.Wrap(err, "failed unmarshalling Win32_Processor")
	}

	return j, nil
}

// Win32Processor data structure
type Win32Processor struct {
	CPUType               string `json:"Name"`
	CurrentProcessorSpeed int    `json:"MaxClockSpeed"`
}
