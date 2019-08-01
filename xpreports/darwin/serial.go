package darwin

import (
	"errors"
	"os/exec"
	"regexp"
)

// GetMachineID generates machine-dependent ID string for a machine.
// All Darwin system should have the IOPlatformSerialNumber attribute.
func GetMachineID() (string, error) {
	out, err := exec.Command("ioreg", "-rd1", "-c", "IOPlatformExpertDevice").Output()
	if err == nil {
		re := regexp.MustCompile("\"IOPlatformSerialNumber\" = \"(.*)\"")
		ret := re.FindStringSubmatch(string(out))
		if len(ret) == 2 {
			return ret[1], nil
		}
	}
	return "", errors.New("can't generate machine ID")
}
