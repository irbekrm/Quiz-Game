package main

import (
  "flag"
  "bufio"
  "strings"
  "os"
  "fmt"
  "encoding/csv"
  "io"
)


var fn = "problems.csv"
var correct int
var incorrect int

func helpPrinter() {
  text := "NAME\n\tQuizz Game\n\nDESCRIPTION\n\tAsks a number of quizz questions. " +
  "Waits for user's answer after each question. Prints the score at the end.\n\n" +
  "-h\n\tprints a short description of the app and then exits\n\n" +
  "-f\n\taccepts name of a file to read quizzes from. It should be a csv file " +
  "in format 'question, answer'. If no filename is passed, reads from the default quizzes file.\n\n"
  fmt.Println(text)
  os.Exit(0)
}

func main() {
  n := flag.String("f", "", "provide file")
  h := flag.Bool("h", false, "print usage description")
  flag.Parse()
  if *h  {
    helpPrinter()
  }
  if *n != "" {
    fn = *n
  }
  reader := bufio.NewReader(os.Stdin)

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
  fmt.Printf("Score: %d/%d\n", correct, correct + incorrect)
}
