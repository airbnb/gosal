// Package xpreports implements cross-platform sal reports.
package xpreports

type Report struct {
	Serial          string `json:"serial"`
	Key             string `json:"key"`
	Name            string `json:"name"`
	DiskSize        string `json:"disk_size"`
	SalVersion      string `json:"sal_version"`
	RunUUID         string `json:"run_uuid"`
	Base64bz2Report string `json:"base_64_bz_2_report"`
}

// Build creates a report for the sal server.
// Build supports darwin, windows and linux and will use
// the appropriate APIs for each system.
func Build() (*Report, error) {

	// buildReport is implented separately for each
	// operating system.
	report, err := buildReport()
	return report, err
}
