package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	// Getting and parsing optional flags from the command line.
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the form of question, answer")
	timeLimit := flag.Int("limit", 30, "the time limit for the quiz in seconds")
	flag.Parse()

	// Opening and parsing CSV file.
	file, err := os.Open(*csvFilename)
	if err != nil {
		exit(fmt.Sprintf("Failed to open the CSV file: %s\n", *csvFilename))
	}
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("Failed to parse the CSV file.")
	}
	problems := parseLines(lines) // Contains each "problem" to be posed to the user.
	correct := 0                  // Counts the score of the user.
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

problemloop:
	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = \n", i+1, p.q)
		answerCh := make(chan string)
		// Spins off the blocking Scanf to a new go routine
		// so that the timer can run uninterrupted in
		// the main routine.
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerCh <- answer
		}()
		// Waits and accepts answers till the timer channel sends a signal.
		select {
		case <-timer.C:
			fmt.Println()
			break problemloop
		case answer := <-answerCh:
			if answer == p.a {
				correct++
			}
		}
	}
	fmt.Printf("You scored %d out of %d.\n", correct, len(problems))
}

// problem holds parsed individual records.
type problem struct {
	q string // Holds the question
	a string // Holds the answer
}

// parseLines takes in a slice containing slices of records
// read from a CSV file by a CSV reader and
// returns a slice containing structs of type problem.
// Each of the structs contain the parsed question and answer strings.
func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			q: line[0],
			a: strings.TrimSpace(line[1]),
		}
	}
	return ret
}

// exit prints the error message to the console and 
//exits application. Used here as shorthand for runtime errors.

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
