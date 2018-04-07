package windows

import (
	"encoding/json"
	"os/exec"
	"strings"

	"github.com/pkg/errors"
)

// GetManagedInstalls is a crappy representation of things installed on a windows client
func get32BitPrograms() ([]Program, error) {
	cmd := exec.Command("powershell", "gp", "HKLM:\\Software\\Microsoft\\Windows\\CurrentVersion\\Uninstall\\*", "|", "ConvertTo-Json")

	// cmd.Stderr = os.Stderr
	o, err := cmd.Output()
	if err != nil {
		return nil, errors.Wrap(err, "exec gp 32 bit programs")
	}

	var p []Program

	if err := json.Unmarshal(o, &p); err != nil {
		return nil, errors.Wrap(err, "failed unmarshalling 32 bit programs")
	}

	return p, nil
}

// GetManagedInstalls is a crappy representation of things installed on a windows client
func get64BitPrograms() ([]Program, error) {
	cmd := exec.Command("powershell", "gp", "HKLM:\\Software\\Wow6432Node\\Microsoft\\Windows\\CurrentVersion\\Uninstall\\*", "|", "ConvertTo-Json")

	// cmd.Stderr = os.Stderr
	o, err := cmd.Output()
	if err != nil {
		return nil, errors.Wrap(err, "exec gp 64 bit programs")
	}

	var p []Program

	if err := json.Unmarshal(o, &p); err != nil {
		return nil, errors.Wrap(err, "failed unmarshalling 64 bit programs")
	}

	return p, nil
}

// marshalManagedInstallFields takes a golang struct and uses field tags to conver to what sal wants
func formatManagedInstallFields(p []Program) ([]byte, error) {

	var all []ManagedInstallsFormatted

	for _, element := range p {
		all = append(all, ManagedInstallsFormatted{
			Description:      element.Description,
			DisplayName:      element.DisplayName,
			Installed:        element.Installed,
			InstalledSize:    element.InstalledSize,
			InstalledVersion: element.InstalledVersion,
			Name:             element.Name,
		})
	}

	data, err := json.MarshalIndent(&all, "", "  ")
	if err != nil {
		return nil, errors.Wrap(err, "message")
	}

	return data, nil
}

// filterManagedInstalls strips out things "apps and programs" wouldn't report to the user (i.e MS Patches, etc)
func filterManagedInstalls(p []Program) ([]Program, error) {

	var filtered []Program

	for _, element := range p {
		if element.DisplayName == "" || strings.Contains(element.Name, "0FF1CE") {
			continue
		} else {
			filtered = append(filtered, Program{
				Description:      element.Description,
				DisplayName:      element.DisplayName,
				Installed:        element.Installed,
				InstalledSize:    element.InstalledSize,
				InstalledVersion: element.InstalledVersion,
				Name:             element.Name,
			})
		}
	}

	return filtered, nil
}

// CreateManagedInstalls returns a formatted and filtered list of apps installed (what you see in programs and features)
func CreateManagedInstalls() (Installs, error) {
	programs64, err := get64BitPrograms()
	if err != nil {
		return nil, errors.Wrap(err, "message")
	}

	programs32, err := get32BitPrograms()
	if err != nil {
		return nil, errors.Wrap(err, "message")
	}

	p := append(programs32[:], programs64[:]...)

	fi, err := filterManagedInstalls(p)
	if err != nil {
		return nil, errors.Wrap(err, "failed to filter programs")
	}

	fo, err := formatManagedInstallFields(fi)
	if err != nil {
		return nil, errors.Wrap(err, "failed to format programs")
	}

	var i Installs

	if err := json.Unmarshal(fo, &i); err != nil {
		return nil, errors.Wrap(err, "failed unmarshalling installs")
	}

	return i, nil
}

// Installs is what is returned to bas64bz2.go
type Installs []map[string]interface{}

// ManagedInstallsFormatted is for remapping the CamelCase keys to snake_case
type ManagedInstallsFormatted struct {
	Description      string `json:"description"`
	DisplayName      string `json:"display_name"`
	Installed        bool   `json:"installed"`
	InstalledSize    int    `json:"installed_size"`
	InstalledVersion string `json:"installed_version"`
	Name             string `json:"name"`
}

// Program data structure that tries to mirror how munki works (but fails mostly)
type Program struct {
	Description      string `json:"Comments"` // MS doesn't carry descriptions, this is close
	DisplayName      string `json:"DisplayName"`
	Installed        bool   // MS ins't munki so this will just default true
	InstalledSize    int    `json:"EstimatedSize"`
	InstalledVersion string `json:"DisplayVersion"`
	Name             string `json:"PSChildName"` // UUID of the items uninstaller
}
