package main

import (
  "github.com/bdemetris/gosal/reports"
  "fmt"
)

func main () {

  test, _ := reports.Get_win32_logicaldisk()

  fmt.Printf("%#v", test)
}
