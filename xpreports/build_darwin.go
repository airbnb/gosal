package xpreports

import (
	"os/exec"

	"github.com/groob/plist"
)

// buildReport creates a report using macOS APIs and paths.
func buildReport() (*Report, error) {
	// some code which would only work on macOS
	// ignore what it actually does, it's just there for demo.

	out, err := exec.Command("/usr/sbin/system_profiler", "SPHardwareDataType", "SPSoftwareDataType", "-xml").CombinedOutput()
	if err != nil {
		return nil, err
	}

	var spData = struct {
		DetailType string `plist:"_dataType"`
	}{}

	if err := plist.Unmarshal(out, &spData); err != nil {
		return nil, err
	}

	report := &Report{Serial: "APPLSERIAL"}
	return report, nil
}
