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

var (
	counter                int
	timeout                bool
	record                 []string
	correct                int
	totalNumberOfQuestions int
)

var fn = "problems.csv"

func shuffle(s [][]string) [][]string {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := range s {
		newPosition := r.Intn(len(s) - 1)
		s[i], s[newPosition] = s[newPosition], s[i]
	}
	return s
}

func sleeper(n int, c chan bool) {
	seconds := time.Duration(n) * time.Second
	time.Sleep(seconds)
	c <- true
}

func helpPrinter() {
	text := "NAME\n\tQuizz Game\n\nDESCRIPTION\n\tAsks a number of quizz questions. " +
		"Waits for user's answer after each question. Prints the score at the end. Default timeout 30s.\n\n" +
		"-h\n\tprints a short usage description and then exits\n\n" +
		"-f\n\taccepts a name of a file to read quizzes from. It should be a csv file " +
		"in format 'question, answer'. If no filename is passed, reads from the default quizzes file.\n\n" +
		"-s\n\tshuffle the questions\n\n" +
		"-t\n\tprovide an int argument X to set the timeout to X seconds (default is 30s)\n\n"
	fmt.Println(text)
	os.Exit(0)
}

func asker(input chan string) {
	for {
		in := bufio.NewReader(os.Stdin)
		result, err := in.ReadString('\n')
		if err != nil {
			fmt.Println("Error", err)
			os.Exit(1)
		}
		result = strings.TrimRight(result, "\n")
		input <- result
	}
}

func main() {
	n := flag.String("f", "", "provide filename")
	h := flag.Bool("h", false, "print usage description")
	s := flag.Bool("s", false, "shuffle the questions")
	t := flag.Int("t", 30, "timeout. requires an int argument")
	flag.Parse()

	if *h {
		helpPrinter()
	}
	if *n != "" {
		fn = *n
	}

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

	input := make(chan string)
	timePassed := make(chan bool)

  in := bufio.NewReader(os.Stdin)
  fmt.Printf("You have %d seconds to answer %d questions. Press Enter when ready\n", *t, totalNumberOfQuestions)
  in.ReadString('\n')
	go sleeper(*t, timePassed)
	go asker(input)

	for counter < len(records) && !timeout {
		record = records[counter]
		fmt.Printf("What is %s?\n", record[0])
		select {
		case i := <-input:
			if i == record[1] {
				correct++
			}
		case <-timePassed:
			fmt.Println("Timeout!")
			timeout = true
		}
		counter++
	}

	fmt.Printf("Score: %d/%d\n", correct, totalNumberOfQuestions)
}
