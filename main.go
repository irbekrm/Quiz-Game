package main

import (
  "os"
  "fmt"
  "encoding/csv"
  "io"
)

func main() {
  fn := "problems.csv"

  f, err := os.Open(fn)
  if err != nil {
    fmt.Println("Error ", err)
    os.Exit(1)
  }

  r := csv.NewReader(f)

  for {
    record, err := r.Read()
    if err == io.EOF {
      break
    }
    if err != nil {
      fmt.Println("Error ", err)
      os.Exit(1)
    }
    fmt.Println(record)
  }
}
