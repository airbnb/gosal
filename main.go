package main

import (
  "os/exec"
  "log"
  "encoding/json"
  "fmt"
)

func main () {

  wb := Get_win32_bios()

  fmt.Printf("%+v\n", wb)
}

func Get_win32_bios() (win32_bios) {
  cmd := exec.Command("powershell", "gwmi", "win32_bios", "|", "convertto-json")

  o, err := cmd.Output()
  if err != nil {
      log.Fatal(err)
  }

  j := win32_bios{}

  json.Unmarshal(o, &j)

  return j
}

type win32_bios struct {
	PSComputerName string  `json:"PSComputerName"`
	SerialNumber   string  `json:"SerialNumber"`
}
