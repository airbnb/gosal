package main

import (
  "os/exec"
  "encoding/json"
  "fmt"
)

func main () {

  test, _ := Get_win32_logicaldisk()

  fmt.Printf("%#v", test)

}

func Get_win32_bios() (win32_bios, error) {
  cmd := exec.Command("powershell", "gwmi", "win32_bios", "|", "convertto-json")

  o, err := cmd.Output()
  if err != nil {
    return win32_bios{}, err
  }

  var j win32_bios

  if err := json.Unmarshal(o, &j); err != nil{
    return win32_bios{}, err
  }

  return j, nil
}

func Get_win32_os() (win32_os, error) {
  cmd := exec.Command("powershell", "gwmi", "Win32_OperatingSystem", "|", "convertto-json")

  o, err := cmd.Output()
  if err != nil {
    return win32_os{}, err
  }

  var j win32_os

  if err := json.Unmarshal(o, &j); err != nil{
    return win32_os{}, err
  }

  return j, nil
}

func Get_win32_logicaldisk() ([]win32_logicaldisk, error) {
  cmd := exec.Command("powershell", "Get-WmiObject", "Win32_LogicalDisk", "|", "convertto-json")

  o, err := cmd.Output()
  if err != nil {
    return nil, err
  }

  var j []win32_logicaldisk

  if err := json.Unmarshal(o, &j); err != nil{
    return nil, err
  }

  return j, nil

}

type win32_logicaldisk struct {
	Name string `json:"Name"`
	Size int64 `json:"Size"`
	FreeSpace int64 `json:"Free"`
}

type win32_bios struct {
	PSComputerName          string   `json:"PSComputerName"`
	SerialNumber            string   `json:"SerialNumber"`
}

type win32_os struct {
  Caption                 string    `json:"Caption"`
  TotalVirtualMemorySize  string    `json:"TotalVirtualMemorySize"`
  TotalVisibleMemorySize  string    `json:"TotalVisibleMemorySize"`
}
