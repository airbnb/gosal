package windows

import (
	"bytes"
	"encoding/base64"
	"strings"
	"time"

	"github.com/airbnb/gosal/config"
	"github.com/airbnb/gosal/xpreports/cm"
	"github.com/dsnet/compress/bzip2"
	"github.com/groob/plist"
	"github.com/pkg/errors"
)

type basereport struct {
	AvailableDiskSpace int
	ConsoleUser        string
	OSFamily           string
	MachineInfo        *MachineInfo
	Facter             cm.Facts
	StartTime          string
}

// BuildBase64bz2Report will return a compressed and encoded string of our report struct
func BuildBase64bz2Report(conf *config.Config) (string, error) {
	var facts map[string]interface{}

	if conf.Management != nil {
		facts, _ = cm.GetFacts(conf.Management.Tool, conf.Management.Path, conf.Management.Command)
	}

	cDrive, err := GetCDrive()
	if err != nil {
		return "", errors.Wrap(err, "bz2: failed getting c: drive")
	}

	machineInfo, err := EmulateMachineInfo()
	if err != nil {
		return "", errors.Wrap(err, "bz2: failed getting machine info")
	}

	computerSystem, err := GetWin32ComputerSystem()
	if err != nil {
		return "", errors.Wrap(err, "bz2: failed getting computer system")
	}

	report := basereport{
		StartTime:          time.Now().Format("01-02-2006"),
		AvailableDiskSpace: cDrive.FreeSpace,
		MachineInfo:        machineInfo,
		ConsoleUser:        strings.Split(computerSystem.UserName, "\\")[1],
		OSFamily:           "Windows",
		Facter:             facts,
	}

	encodedReport, err := report.CompressAndEncode()
	if err != nil {
		return "", errors.Wrap(err, "bz2: failed to compress and encode report")
	}

	return encodedReport, nil
}

func (r *basereport) CompressAndEncode() (string, error) {

	var buf bytes.Buffer

	bzw, err := bzip2.NewWriter(&buf, &bzip2.WriterConfig{Level: bzip2.BestSpeed})
	if err != nil {
		return "", errors.Wrap(err, "bz2: failed to bzip2")
	}
	defer bzw.Close()

	enc := plist.NewEncoder(bzw)
	enc.Indent("  ")

	if err = enc.Encode(r); err != nil {
		return "", errors.Wrap(err, "bz2: failed to encode plist")
	}
	bzw.Close()

	report := base64.StdEncoding.EncodeToString(buf.Bytes())
	if err != nil {
		return "", errors.Wrap(err, "bz2: failed to base64 encode the report string")
	}

	return report, nil
}
