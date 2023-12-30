package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	flag.Parse()

	file, err := os.Open(*csvFilename)
	if err != nil {
		exit(fmt.Sprintf("Failed to open CSV file: %s", *csvFilename))
	}
	defer file.Close()

	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("Failed to parse the provided CSV file")
	}

	var correct int
	for _, line := range lines {
		question := line[0]
		answer := strings.TrimSpace(line[1])

		fmt.Printf("Question: %s\n", question)

		var userAnswer string
		fmt.Print("Your answer: ")
		fmt.Scanln(&userAnswer)

		if userAnswer == answer {
			correct++
		}
	}

	fmt.Printf("\nYou got %d out of %d questions correct.\n", correct, len(lines))
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
