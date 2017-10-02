package reports

import (
	"encoding/json"
	"fmt"
	"log"
)

// EmulateMachineInfo copies its behavior from macOS, and provides struct data to Sal
func EmulateMachineInfo() (MachineInfo, error) {

	win32OS, err := GetWin32OS()
	if err != nil {
		// TODO return the error here?
		log.Printf("reports: getting win32 os: %s", err)
	}

	systemProfile, err := EmulateSystemProfile()
	if err != nil {
		// TODO return the error here?
		log.Printf("reports: system profile failed: %s", err)
	}

	fmt.Println(systemProfile)

	report := MachineInfo{
		OSVers:        win32OS.Caption,
		SystemProfile: systemProfile,
	}

	return report, nil
}

// MachineInfo is a plist item that sal expects to parse relative data
type MachineInfo struct {
	OSVers        string
	SystemProfile SystemProfile
}

type SystemProfile []struct {
	DataType string `json:"_dataType"`
	Items    *Items `json:"_items"`
}

type Items []struct {
	MachineModel          string `json:"machine_model"`
	CPUType               string `json:"cpu_type"`
	CurrentProcessorSpeed string `json:"current_processor_speed"`
	PhysicalMemory        string `json:"physical_memory"`
}

// EmulateSystemProfile creates the necessary system_profile
func EmulateSystemProfile() (SystemProfile, error) {

	values := `[{"_dataType": "SPHardwareDataType","_items": [{"machine_model": "hi","cpu_type": "hello","current_processor_speed": "hi","physical_memory": "hello"}]}]`

	var i SystemProfile

	if err := json.Unmarshal([]byte(values), &i); err != nil {
		log.Printf("reports: unmarshaling system profile: %s", err)
	}

	fmt.Printf("%+v", i)

	return i, nil

}
