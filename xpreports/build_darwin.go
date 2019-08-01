package xpreports

import (
	"strconv"

	"github.com/airbnb/gosal/config"
	"github.com/airbnb/gosal/version"
	"github.com/airbnb/gosal/xpreports/darwin"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"github.com/shirou/gopsutil/host"
)

// buildReport creates a report using linux APIs and paths.
func buildReport(conf *config.Config) (*Report, error) {
	u1 := uuid.NewV4().String()

	h, _ := host.Info()

	s, _ := darwin.GetMachineID()

	disk, err := darwin.GetDisk()
	if err != nil {
		return nil, errors.Wrap(err, "reports: getting disk info")
	}

	encodedCompressedPlist, err := darwin.BuildBase64bz2Report(conf)
	if err != nil {
		return nil, errors.Wrap(err, "reports: getting plist")
	}

	// Get version information
	v := version.Version()

	report := &Report{
		Serial:          s,
		Key:             conf.Key,
		Name:            h.Hostname,
		DiskSize:        strconv.Itoa(disk.Size),
		SalVersion:      v.Version,
		RunUUID:         u1,
		Base64bz2Report: encodedCompressedPlist,
	}

	// fmt.Printf("%+v\n", report)
	return report, nil
}
