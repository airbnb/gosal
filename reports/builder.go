package reports

import (
	"log"
	"strconv"
	"strings"
	"fmt"
	"bytes"
	"encoding/base64"

	"github.com/satori/go.uuid"
	"github.com/groob/plist"
	"github.com/dsnet/compress/bzip2"
)

// BuildReport builds the report object
func BuildReport(apiKey string) Report {

	win32Bios, _ := GetWin32Bios()
	CDrive, err := GetCDrive()
	if err != nil {
		// TODO return the error here?
		log.Printf("reports: getting win32 disk: %s", err)
	}

	win32ComputerSystem, err := GetWin32ComputerSystem()
	if err != nil {
		// TODO return the error here?
		log.Printf("reports: getting win32 computer system: %s", err)
	}
	u1 := uuid.NewV4().String()

	// This is our report debug
	base, _ := BuildBase64bz2Report()

	// fmt.Printf("%+v\n", base)
	var buf bytes.Buffer

	bzw, err := bzip2.NewWriter(&buf, &bzip2.WriterConfig{Level: bzip2.BestSpeed})
    if err != nil {
        log.Fatal(err)
    }
    defer bzw.Close()

		enc := plist.NewEncoder(bzw)
    enc.Indent("  ")

		if err := enc.Encode(base); err != nil {
        log.Fatal(err)
    }
		bzw.Close()

	report := Report{
		Serial:     win32Bios.SerialNumber,
		Key:        apiKey,
		Name:       win32Bios.PSComputerName,
		DiskSize:   strconv.Itoa(CDrive.Size),
		SalVersion: strconv.Itoa(1),
		RunUUID:    u1,
		UserName:   strings.Split(win32ComputerSystem.UserName, "\\")[1],
		Base64bz2Report: base64.StdEncoding.EncodeToString(buf.Bytes()),
	}

	fmt.Printf("%+v\n", report)
	return report
}

// Report structure
type Report struct {
	Serial     string
	Key        string
	Name       string
	DiskSize   string
	SalVersion string
	RunUUID    string
	UserName   string
	Base64bz2Report string
}
