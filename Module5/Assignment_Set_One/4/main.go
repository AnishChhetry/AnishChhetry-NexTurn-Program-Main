package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Question struct {
	QuestionText string
	Options      [4]string
	CorrectAns   int
}

var questionBank = []Question{
	{
		QuestionText: "What is 2x7?",
		Options:      [4]string{"1. 1", "2. 23", "3. 32", "4. 14"},
		CorrectAns:   4,
	},
	{
		QuestionText: "What is the capital of India?",
		Options:      [4]string{"1. Delhi", "2. Aizawl", "3. Hyderabad", "4. Raipur"},
		CorrectAns:   1,
	},
	{
		QuestionText: "Find the odd one:",
		Options:      [4]string{"1. Cake", "2. Kheer", "3. Noodles", "4. Laddoo"},
		CorrectAns:   3,
	},
}

const questionTimeLimit = 30 * time.Second

func TakeQuiz() {
	var score int
	reader := bufio.NewReader(os.Stdin)

	for i, question := range questionBank {
		fmt.Printf("Question %d: %s\n", i+1, question.QuestionText)
		for _, option := range question.Options {
			fmt.Println(option)
		}

		answerChan := make(chan int)
		timer := time.NewTimer(questionTimeLimit)

		go func() {
			for {
				fmt.Print("Enter your answer (or type 'exit' to quit): ")
				input, _ := reader.ReadString('\n')
				input = strings.TrimSpace(input)

				if input == "exit" {
					close(answerChan)
					return
				}

				answer, err := strconv.Atoi(input)
				if err != nil || answer < 1 || answer > 4 {
					fmt.Println("Invalid input. Please enter a number between 1 and 4.")
					continue
				}
				answerChan <- answer
				return
			}
		}()

		select {
		case answer, ok := <-answerChan:
			if !ok {
				fmt.Println("Exiting quiz.")
				return
			}
			if answer == question.CorrectAns {
				fmt.Println("Correct!")
				score++
			} else {
				fmt.Println("Wrong answer.")
			}
		case <-timer.C:
			fmt.Println("Time's up for this question!")
		}
		fmt.Println()
	}

	fmt.Printf("Quiz completed! Your score is: %d/%d\n", score, len(questionBank))
	if score == len(questionBank) {
		fmt.Println("Performance: Excellent")
	} else if score >= len(questionBank)/2 {
		fmt.Println("Performance: Good")
	} else {
		fmt.Println("Performance: Needs Improvement")
	}
}

func main() {
	fmt.Println("Welcome to the Online Examination System!")
	fmt.Println("You will have 30 seconds to answer each question. Type 'exit' to quit early.")
	fmt.Println("-----------------------------------------------------------")

	TakeQuiz()
}
