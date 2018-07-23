// Quizzes the user based on a csv file with 2 fields on each record(line)
// containing a problem and an answer.

package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type problem struct {
	statement string // User prompt
	answer    string // Correct Responce
}

func main() {
	var filename string

	// Uses 'problems.csv' if no argument is passed
	if len(os.Args) == 1 {
		filename = "problems.csv"
		fmt.Println("No filename specified. Using 'problems.csv' for config file...")
	} else {
		filename = os.Args[1]
	}

	problems := readFromFile(filename)
	r := bufio.NewReader(os.Stdin)
	var right int

	fmt.Print("\nQUIZ BEGIN\n\n")
	for _, problem := range problems {
		fmt.Print(problem.statement, " ")
		guess, _ := r.ReadString('\n')

		// The delimiter is recorded with the data, so it should be removed
		// in order to ensure proper comparisons to what the user had in mind
		if strings.TrimRight(guess, "\n") == problem.answer {
			right++
		}
	}

	fmt.Println("\nCorrect: ", right)
	fmt.Println("Total: ", len(problems))
	fmt.Print("\nQUIZ END\n\n")

}

func readFromFile(filename string) []problem {
	cf, eo := os.Open(filename)

	if eo != nil {
		log.Fatal("Error: ", eo)
	}

	cr := csv.NewReader(cf)

	// The csv config file should have exactly 2 fields for all records-
	// anything else is considerd a corruption and is treated as a fatal error.
	cr.FieldsPerRecord = 2
	var ps []problem

	for {
		// Get a string slice from each column
		record, err := cr.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal("Error: ", err)
		}

		ps = append(ps, newProb(record))
	}

	return ps
}

func newProb(r []string) problem {
	return problem{statement: r[0], answer: strings.TrimRight(r[1], "\r\n")}
}
