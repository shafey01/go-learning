package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

const file = "problems.csv"

var result = 0

func main() {
	// Open file from system
	f, err := os.Open(file)
	if err != nil {
		fmt.Printf("Error from read operation OS %v \n", err)

		return
	}

	// Defer to close the file in the end of program
	defer f.Close()

	// Read file from the reference when we open it
	reader := csv.NewReader(f)
	questions, err := reader.ReadAll()
	if err != nil {
		fmt.Printf("Error from read operation %v \n", err)
	}

	startQuiz(questions)
	fmt.Printf("result is: %d from %d \n", result, len(questions))
}

// function to start Quiz and return the result in intger
func startQuiz(questions [][]string) {
	for _, record := range questions {
		q, answer := record[0], record[1]
		userAnswer := ""
		fmt.Println(q)
		_, err := fmt.Scanln(&userAnswer)
		if err != nil {
			fmt.Printf("Error from read answer operation OS %v \n", err)

			return
		}
		if userAnswer == answer {
			result++
		}

	}
}
