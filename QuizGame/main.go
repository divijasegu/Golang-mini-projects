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
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	timeLimit := flag.Int("limit", 30, "the time limit for the quiz in seconds")
	flag.Parse()

	file, err := os.Open(*csvFilename)

	if err != nil {
		exit(fmt.Sprintf("Failed to open the CSV file: %s\n", *csvFilename))
	}
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("Failed to parse the provided CSV file.")
	}
	problems := parseLines(lines)
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)
	<-timer.C

	correct := 0
	for i, p := range problems {
		select {
		case <-timer.C:
			fmt.Printf("You scored %d out of %d.\n", correct, len(problems))
			return
		default:
			fmt.Printf("Problem #%d: %s = \n", i+1, p.question)
			var answer string
			fmt.Scanf("%s\n", &answer)
			if answer == p.answer {
				fmt.Println("Correct!")
				correct++
			} else {
				fmt.Println("Wrong!")
			}

		}

	}

	fmt.Printf("You scored %d out of %d.\n", correct, len(problems))

}

type problem struct {
	question string
	answer   string
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			question: line[0],
			answer:   strings.TrimSpace(line[1]),
		}
	}

	return ret

}
