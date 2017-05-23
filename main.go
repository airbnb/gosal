package main

import (
  "os/exec"
  "fmt"
  "log"
  "bufio"
  "strings"
)

func main () {
  cmd := exec.Command("wmic", "bios", "list", "full", "/format:csv")

  o, err := cmd.Output()
  if err != nil {
      log.Fatal(err)
  }

  s := string(o)

  var csv_lines []string

  scanner := bufio.NewScanner(strings.NewReader(s))
  for scanner.Scan() {
    csv_lines = append(csv_lines, scanner.Text())
  }

  keys := csv_lines[1]
  values := csv_lines[2]

  fmt.Println(keys + "\n")
  fmt.Println(values)


  // r := csv.NewReader(strings.NewReader(s))
  // fmt.Println(r)

  // for {
  //   record, err := r.ReadAll()
  //   if err == io.EOF {
  //     break
  //   }
  //   if err != nil {
  //     log.Fatal(err)
  //     fmt.Println(record)
  //   }

  //   fmt.Println(record)
  // }



  // fmt.Println(s)
}
