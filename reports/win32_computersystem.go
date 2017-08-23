package reports

import (
	"encoding/json"
	"os/exec"
)

func Get_win32_computersystem() (win32_computersystem, error) {
	cmd := exec.Command("powershell", "gwmi", "win32_ComputerSystem", "|", "ConvertTo-Json")

	o, err := cmd.Output()
	if err != nil {
		return win32_computersystem{}, err
	}

	var j win32_computersystem

	if err := json.Unmarshal(o, &j); err != nil {
		return win32_computersystem{}, err
	}

	return j, nil
}

type win32_computersystem struct {
	UserName string `json:"UserName"`
}
