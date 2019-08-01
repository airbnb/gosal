package linux

import (
	"os/exec"
	"strings"
)

func LoggedInUsers() ([]string, error) {

	cmd := exec.Command("who", "-us")
	users, _ := cmd.CombinedOutput()

	cleaned := []string{}

	for _, user := range strings.Split(string(users), "\n") {
		clean := true
		col := strings.Split(user, " ")

		if len(col) > 0 {
			for _, cleanedU := range cleaned {
				u := strings.TrimSpace(col[0])
				if len(u) == 0 || strings.Compare(cleanedU, col[0]) == 0 {
					clean = false
				}
			}
			if clean {
				cleaned = append(cleaned, col[0])
			}
		}
	}

	return cleaned, nil
}
