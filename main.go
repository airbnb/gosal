package main

import (
  "github.com/bdemetris/gosal/reports"
  "fmt"
  "encoding/json"
)

func main () {

  win32_os, _ := reports.Get_win32_os()
  win32_bios, _ := reports.Get_win32_bios()
  win32_logicaldisk, _ := reports.Get_win32_logicaldisk()

  var report map[string]string

  json_win32_os, _ := json.Marshal(win32_os)
  json.Unmarshal(json_win32_os, &report)

  json_win32_bios, _ := json.Marshal(win32_bios)
  json.Unmarshal(json_win32_bios, &report)

  json_win32_logicaldisk, _ := json.Marshal(win32_logicaldisk)
  json.Unmarshal(json_win32_logicaldisk, &report)

  json_report, _ := json.Marshal(report)
  fmt.Println(string(json_report))
  // test, _ := reports.Get_win32_logicaldisk()
  //
  // fmt.Printf("%#v", test)
}
