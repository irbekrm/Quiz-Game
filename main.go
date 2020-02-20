package main

//importing dependencies
import (
	"bufio"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
    "strconv"
)

//declaring variables
var (
	counter                int
	timeout                bool
	correct                int
)

//defining function which times the user
func sleeper(n int, c chan bool) {
	seconds := time.Duration(n) * time.Second
	time.Sleep(seconds)
	c <- true
}

//defining a function which prints help messages for the user
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

//definging a function to read the user input
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

//defining a main function
func main() {
	h := flag.Bool("h", false, "print usage description")
	t := flag.Int("t", 30, "timeout. requires an int argument")
	flag.Parse()

	if *h {
		helpPrinter()
	}

	input := make(chan string)
	timePassed := make(chan bool)

	in := bufio.NewReader(os.Stdin)
	fmt.Printf("You have %d seconds to answer 12 questions. Press Enter when ready\n", *t)
	in.ReadString('\n')
	go sleeper(*t, timePassed)
	go asker(input)

    //my contribution lies here
    //the program now generates two random numbers, calculates the answer and compares 
    //it to the users answer
	for counter < 12 && !timeout {
        rand.Seed(time.Now().UTC().UnixNano())
		var int1 int = rand.Intn(10)
		var int2 int = rand.Intn(10)
		var answer int = int1 + int2
		var newAnswer string = strconv.Itoa(answer)
		fmt.Printf("What is %d+%d?\n", int1, int2)
		select {
		case i := <-input:
			if i == newAnswer {
				correct++
			}
		case <-timePassed:
			fmt.Println("Timeout!")
			timeout = true
		}
		counter++
	}

	fmt.Printf("Score: %d/12\n", correct)
}

