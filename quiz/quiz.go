package quiz

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"time"
)

func ExerciseOne(timeLimit time.Duration) {

	totalCorrect := 0
	totalLines := 0

	file, fileError := os.Open("./quiz/problems.csv")
	defer file.Close()
	if fileError != nil {
		return
	}
	fileLines, err := csv.NewReader(file).ReadAll()
	if err != nil {
		return
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Please press enter to begin.")
	_, _ = reader.ReadString('\n')
	if timeLimit <= 0 {
		timeLimit = 30
	}
	timer := time.NewTimer(timeLimit * time.Second)
	// for each line, break the first and second parts, trimming out the words
	// iterate through the file lines
	for _, lines := range fileLines {
		totalLines += 1
		question := lines[0]
		answer := lines[1]

		answerCh := make(chan string)

		go func() {
			var userAnswer string
			fmt.Printf("%s = ", question)
			fmt.Scanln(&userAnswer)
			answerCh <- userAnswer
		}()

		select {
		case <-timer.C:
			fmt.Println("\nTime's up!")
			fmt.Printf("Total correct: %d, Total lines: %d\n", totalCorrect, totalLines)
			os.Exit(0)
		case userAnswer := <-answerCh:
			if userAnswer == answer {
				totalCorrect += 1
			}
		}
	}

	fmt.Printf("Total correct: %d, Total lines: %d\n", totalCorrect, totalLines)
	return
}
