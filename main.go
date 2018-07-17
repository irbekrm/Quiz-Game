package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

var fn = "problems.csv"
var correct int
var totalNumberOfQuestions int

func shuffle(s [][]string) [][]string {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := range s {
		newPosition := r.Intn(len(s) - 1)
		s[i], s[newPosition] = s[newPosition], s[i]
	}
	return s
}

func sleeper(n int){
  seconds := time.Duration(n) * 30 * time.Second
  time.Sleep(seconds)
  fmt.Println("Timeout!")
  os.Exit(0)
}

func helpPrinter() {
	text := "NAME\n\tQuizz Game\n\nDESCRIPTION\n\tAsks a number of quizz questions. " +
		"Waits for user's answer after each question. Prints the score at the end.\n\n" +
		"-h\n\tprints a short usage description and then exits\n\n" +
		"-f\n\taccepts name of a file to read quizzes from. It should be a csv file " +
		"in format 'question, answer'. If no filename is passed, reads from the default quizzes file.\n\n" +
		"-s\n\tshuffle the questions\n\n" +
    "-t\n\tprovide an int argument X to increase the default timeout value (30s) to X * 30s.\n\n"
	fmt.Println(text)
	os.Exit(0)
}

func main() {
	n := flag.String("f", "", "provide filename")
	h := flag.Bool("h", false, "print usage description")
	s := flag.Bool("s", false, "shuffle the questions")
  t := flag.Int("t", 1, "timeout. requires an int argument")
	flag.Parse()

	if *h {
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

	records, err := r.ReadAll()

	if err != nil {
		fmt.Println("Error", err)
	}

	totalNumberOfQuestions = len(records)

	if *s {
		records = shuffle(records)
	}

  go sleeper(*t)

	for _, record := range records {
		fmt.Println("What is " + record[0] + "?")
		input, _ := reader.ReadString('\n')
		input = strings.TrimRight(input, "\n")
		if input == record[1] {
			correct++
		}
	}
	fmt.Printf("Score: %d/%d\n", correct, totalNumberOfQuestions)
}
