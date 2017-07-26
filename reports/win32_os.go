package reports

import (
  "os/exec"
  "encoding/json"
  "fmt"
)

func Get_win32_os() (win32_os, error) {
  cmd := exec.Command("powershell", "gwmi", "Win32_OperatingSystem", "|", "convertto-json")

  o, err := cmd.Output()

  if err != nil {
    fmt.Println("error here")
    return win32_os{}, err
  }

  var j win32_os

  fmt.Println(j)

  if err := json.Unmarshal(o, &j); err != nil{
    return win32_os{}, err
  }

  return j, nil
}

type win32_os struct {
  Caption                 string    `json:"Caption"` //os version
  TotalVirtualMemorySize  int    `json:"TotalVirtualMemorySize"`
  TotalVisibleMemorySize  int    `json:"TotalVisibleMemorySize"`
}
