package main

import (
  "os/exec"
  "fmt"
  "log"
  "strings"
)

func main () {
  echo := exec.Command("wmic", "bios", "list", "full")

  b, err := echo.Output()
  if err != nil {
      log.Fatal(err)
  }

  s := string(b)

  fmt.Println((strings.Split(s, "\r"))[1])
}
