package main

import (
  "bufio"
  "strings"
  "os"
  "fmt"
  "encoding/csv"
  "io"
)


var correct int
var incorrect int

func main() {
  reader := bufio.NewReader(os.Stdin)
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
    fmt.Println("What is " + record[0] + "?")
    input, _ := reader.ReadString('\n')
    input = strings.TrimRight(input, "\n")
    if input == record[1] {
      correct++
    } else {
      incorrect++
    }
  }
  fmt.Println("Correct:", correct, "incorrect:", incorrect)
  
}
