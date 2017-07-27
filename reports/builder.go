package reports

import (
	"github.com/salopensource/gosal/utils"
	"github.com/satori/go.uuid"
	"strconv"
)

// build the report object
func BuildReport() Report {

	win32_bios, _ := Get_win32_bios()
	win32_logicaldisk, _ := Get_win32_logicaldisk()
	u1 := uuid.NewV4().String()

	report := Report{
		Serial:     win32_bios.SerialNumber,
		Key:        utils.LoadConfig("./config.json").Key,
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
