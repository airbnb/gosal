package windows

import (
	"encoding/json"
	"os/exec"

	"github.com/pkg/errors"
)

// GetManagedInstalls is a crappy representation of things installed on a windows client
func GetManagedInstalls() ([]ManagedInstalls, error) {
	cmd := exec.Command("powershell", "gp", "HKLM:\\Software\\Microsoft\\Windows\\CurrentVersion\\Uninstall\\*", "|", "ConvertTo-Json")

	// cmd.Stderr = os.Stderr
	o, err := cmd.Output()
	if err != nil {
		return nil, errors.Wrap(err, "exec gp")
	}

	var j []ManagedInstalls

	if err := json.Unmarshal(o, &j); err != nil {
		return nil, errors.Wrap(err, "failed unmarshalling Installs")
	}

	return j, nil
}

// marshalManagedInstallFields takes a golang struct and uses field tags to conver to what sal wants
func MarshalManagedInstallFields() ([]ManagedInstallsFormatted, error) {

	mi, err := GetManagedInstalls()
	if err != nil {
		return nil, errors.Wrap(err, "message")
	}

	for _, element := range mi {
		json.Marshal(&ManagedInstallsFormatted{
			Description:      element.Description,
			DisplayName:      element.DisplayName,
			Installed:        element.Installed,
			InstalledSize:    element.InstalledSize,
			InstalledVersion: element.InstalledVersion,
			Name:             element.Name,
		})

	}

	// fmt.Println(string(mif))

	return nil, nil
}

type ManagedInstallsFormatted struct {
	Description      string `json:"description"`
	DisplayName      string `json:"display_name"`
	Installed        bool   `json:"installed"`
	InstalledSize    int    `json:"installed_size"`
	InstalledVersion string `json:"installed_version"`
	Name             string `json:"name"`
}

// ManagedInstalls array of dicts
//   description string
//   display_name string
//   installed bool
//   installed_size int
//   installed_version string
//   name string

// ManagedInstalls data structure that tries to mirror how munki works (but fails mostly)
type ManagedInstalls struct {
	Description      string `json:"Comments"` // MS doesn't carry descriptions, this is close
	DisplayName      string `json:"DisplayName"`
	Installed        bool   // MS ins't munki so this will just default true
	InstalledSize    int    `json:"EstimatedSize"`
	InstalledVersion string `json:"DisplayVersion"`
	Name             string `json:"PSChildName"` // UUID of the installed item
}
