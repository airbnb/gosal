package reports

import (
	"log"
	"strconv"
	"strings"

	"github.com/satori/go.uuid"
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

	report := Report{
		Serial:     win32Bios.SerialNumber,
		Key:        apiKey,
		Name:       win32Bios.PSComputerName,
		DiskSize:   strconv.Itoa(CDrive.Size),
		SalVersion: strconv.Itoa(1),
		RunUUID:    u1,
		UserName:   strings.Split(win32ComputerSystem.UserName, "\\")[1],
	}

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
}
