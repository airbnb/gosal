package reports

import (
  "fmt"
  "github.com/satori/go.uuid"
  "github.com/bdemetris/gosal/utils"
)

// build the report object
func BuildReport() (Report){

  win32_bios, _ := Get_win32_bios()
  win32_logicaldisk, _ := Get_win32_logicaldisk()
  u1 := uuid.NewV4().String()

  report := Report{
    serial:       win32_bios.SerialNumber,
    key:          utils.LoadConfig("./config.json").Key,
    name:         win32_bios.PSComputerName,
    disk_size:    win32_logicaldisk[1].Size,
    sal_version:  1,
    run_uuid:     u1,
  }
  fmt.Println(report)
  return report
}

// report structure
type Report struct {
  serial          string
  key             string
  name            string
  disk_size        int64
  sal_version      int64
  run_uuid         string
}
