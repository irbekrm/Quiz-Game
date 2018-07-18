Quiz-Game
===
The first exercise from https://gophercises.com/exercises/

### Description

A timed quiz. Asks user a number of questions and prints the score at the end.

The default version has a timeout of 30s and the quizzes are read from the *problems.csv* file included.

User can set a different timeout, provide another quiz file and shuffle quiz questions.

### Use

(If you don't have *go* installed and go workspace set up, follow the instructions here: https://golang.org/doc/install)

*go get github.com/irbekrm/Quiz-Game* download and install this package

*cd ~/go/src/github.com/irbekrm/Quiz-Game* go to the package directory

*go run ./main.go -h* print help

*go run ./main.go* run the default version

*go run ./main.go -s* shuffle the questions

*go run ./main.go -f \<filename\>* provide own csv file with quizzes. File should be in format *question,answer*

*go run ./main.go -t \<x\>* set timeout to x seconds

### Implementation

The main challenge for me was to track the time. If the time runs out whilst the program is waiting for the user input
for a quiz question, the program should stop waiting for the user input, print the score and exit.

To solve this, I created two go subroutines, one runs the timer, the other waits for user input. Each has an associated channel-
the timeout subroutine sends *true* when the time runs out and the user input subroutine sends the user's answer,
when it recieves one. When the user is ready,
the main process spins off both subroutines. It then runs a *for* loop for as long as a counter variable that I use to
loop through the questions in the quiz list is less than the length of the list and a *timeout* variable is *false*. Inside the loop
there is a *select* statement that checks for data being sent from either channel. If a user's answer is sent,
that will be recorded and the counter variable incremented. If the timer subroutine sends *true*, the *timeout* variable will be set to *true*,
which will cause the *for* loop to exit.
