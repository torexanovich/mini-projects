package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
	"math/rand"
)

func main() {
	csvFileName := flag.String("csv", "problems.csv", "CSV file with 'question,answer'")
	timeLimit := flag.Int("time", 30, "time limit for the quiz in seconds")
	shuffle := flag.Bool("shuffle", false, "shuffle quiz order")
	flag.Parse()

	questions := readCSV(*csvFileName)
	if *shuffle {
		shuffleQuestions(questions)
	}

	fmt.Println("Press Enter to start the quiz...")
	fmt.Scanln()

	correct := runQuiz(questions, *timeLimit)

	fmt.Printf("\nYou got %d out of %d questions correct.\n", correct, len(questions))
}

func readCSV(filename string) [][]string {
	file, err := os.Open(filename)
	if err != nil {
		exit(fmt.Sprintf("Failed to open CSV file: %s", filename))
	}
	defer file.Close()

	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("Failed to parse the provided CSV file.")
	}

	return lines
}

func shuffleQuestions(questions [][]string) {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(questions), func(i, j int) {
		questions[i], questions[j] = questions[j], questions[i]
	})
}

func runQuiz(questions [][]string, timeLimit int) int {
	timer := time.NewTimer(time.Duration(timeLimit) * time.Second)
	correct := 0

questionLoop:
	for _, q := range questions {
		question, answer := q[0], strings.TrimSpace(q[1])

		fmt.Printf("\nQuestion: %s\nYour answer: ", question)

		select {
		case <-timer.C:
			fmt.Println("\nTime's up! Quiz completed.")
			break questionLoop
		case userAnswer := <-getUserAnswer():
			if strings.TrimSpace(userAnswer) == answer {
				correct++
			}
		}
	}

	return correct
}

func getUserAnswer() <-chan string {
	answerCh := make(chan string)
	go func() {
		var userAnswer string
		fmt.Scanln(&userAnswer)
		answerCh <- userAnswer
	}()
	return answerCh
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
