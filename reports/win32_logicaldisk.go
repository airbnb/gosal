package reports

import (
  "os/exec"
  "encoding/json"
)

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
	Name       string  `json:"Name"`
	Size       int64   `json:"Size"`
	FreeSpace  int64   `json:"Free"`
}
