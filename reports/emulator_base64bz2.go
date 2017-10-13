package reports

import (
	"bytes"
	"encoding/base64"
	"log"
	"strings"

	"github.com/dsnet/compress/bzip2"
	"github.com/groob/plist"
)

type basereport struct {
	AvailableDiskSpace int
	ConsoleUser        string
	OSFamily           string
	MachineInfo        MachineInfo
}

// BuildBase64bz2Report will return a compressed and encoded string of our report struct
func BuildBase64bz2Report() (string, error) {

	cDrive, err := GetCDrive()
	if err != nil {
		// TODO return the error here?
		log.Printf("reports: getting win32 disk: %s", err)
	}

	machineInfo, err := EmulateMachineInfo()
	if err != nil {
		// TODO return the error here?
		log.Printf("reports: problem with emulating machine info: %s", err)
	}

	computerSystem, err := GetWin32ComputerSystem()
	if err != nil {
		// TODO return the error here?
		log.Printf("machine info: computer system information failed: %s", err)
	}

	report := basereport{
		AvailableDiskSpace: cDrive.FreeSpace,
		MachineInfo:        machineInfo,
		ConsoleUser:        strings.Split(computerSystem.UserName, "\\")[1],
		OSFamily:           "Windows",
	}

	// fmt.Println(report)

	encodedReport, err := report.CompressAndEncode()
	if err != nil {
		// TODO return the error here?
		log.Printf("compress and encode failed: %s", err)
	}

	return encodedReport, nil
}

func (r *basereport) CompressAndEncode() (string, error) {

	var buf bytes.Buffer

	bzw, err := bzip2.NewWriter(&buf, &bzip2.WriterConfig{Level: bzip2.BestSpeed})
	if err != nil {
		log.Fatal(err)
	}
	defer bzw.Close()

	enc := plist.NewEncoder(bzw)
	enc.Indent("  ")

	if err = enc.Encode(r); err != nil {
		log.Fatal(err)
	}
	bzw.Close()

	report := base64.StdEncoding.EncodeToString(buf.Bytes())
	if err != nil {
		log.Fatal(err)
	}

	return report, nil
}
