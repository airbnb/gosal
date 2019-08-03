package darwin

import (
	"os/exec"
	"regexp"
)

// GetMacOSComputerSystem exports  powershell class
func GetMacOSComputerSystem() (MacOSComputerSystem, error) {
	out, err := exec.Command("ioreg", "-rd1", "-c", "IOPlatformExpertDevice").Output()
	if err == nil {
		re := regexp.MustCompile("\"model\" = \"(.*)\"")
		ret := re.FindStringSubmatch(string(out))
		if len(ret) == 2 {
			// return ret[1], nil
		}
	}
	//	return "", errors.New("can't generate machine ID")

	compSys := MacOSComputerSystem{
		UserName:     "gavin",
		Manufacturer: "apple",
		Model:        "Precision 5820 Tower",
	}

	return compSys, nil
}

// Win32ComputerSystem structure
type MacOSComputerSystem struct {
	UserName     string
	Manufacturer string
	Model        string
}
