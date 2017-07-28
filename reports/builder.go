package reports

import (
	"log"
	"strconv"

	"github.com/satori/go.uuid"
)

// build the report object
func BuildReport(apiKey string) Report {

	win32_bios, _ := Get_win32_bios()
	win32_logicaldisk, err := Get_win32_logicaldisk()
	if err != nil {
		// TODO return the error here?
		log.Printf("reports: getting win32 disk: %s", err)
	}
	u1 := uuid.NewV4().String()

	report := Report{
		Serial:     win32_bios.SerialNumber,
		Key:        apiKey,
		Name:       win32_bios.PSComputerName,
		DiskSize:   strconv.Itoa(win32_logicaldisk[1].Size),
		SalVersion: strconv.Itoa(1),
		RunUUID:    u1,
	}

	return report
}

// report structure
type Report struct {
	Serial     string
	Key        string
	Name       string
	DiskSize   string
	SalVersion string
	RunUUID    string
}
