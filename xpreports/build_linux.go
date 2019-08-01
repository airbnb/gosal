package xpreports

import (
	"github.com/airbnb/gosal/config"
	"github.com/airbnb/gosal/version"
	"github.com/airbnb/gosal/xpreports/linux"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
)

// buildReport creates a report using linux APIs and paths.
func buildReport(conf *config.Config) (*Report, error) {
	u1 := uuid.NewV4().String()

	encodedCompressedPlist, err := linux.BuildBase64bz2Report(conf)
	if err != nil {
		return nil, errors.Wrap(err, "reports: getting plist")
	}

	// Get version information
	v := version.Version()

	report := &Report{
		Serial:          "123456789",
		Key:             conf.Key,
		Name:            "PC-Name",
		DiskSize:        "40",
		SalVersion:      v.Version,
		RunUUID:         u1,
		Base64bz2Report: encodedCompressedPlist,
	}

	// fmt.Printf("%+v\n", report)
	return report, nil
}
