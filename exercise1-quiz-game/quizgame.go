package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"time"
)

type QuizData struct {
	Question map[string]string
}

func NewQuizData() *QuizData {
	return &QuizData{
		Question: map[string]string{},
	}
}

func (q *QuizData) quizReader(reader *csv.Reader) {

	buff, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	for _, question := range buff {
		q.putNewQuestion(question)
	}
}

func (q *QuizData) putNewQuestion(line []string) {
	question, answer := line[0], line[1]
	q.Question[question] = answer
}

func askQuestions(q *QuizData) int {
	correctAnswers := 0
	index := 1
	ch := make(chan bool)
	timer := time.NewTimer(time.Second * 2)

	go func() {
		for question, answer := range q.Question {
			var a string 
			fmt.Printf("%d. %s\n", index, question)
			fmt.Scanln(&a)
			if a == answer {
				ch <-true
			} else {
				ch <-false
			}
			index++
		}
	}()

	for {
		select {
		case isCorrect := <-ch:
			if isCorrect {
				correctAnswers++
			}
		case <-timer.C:
			log.Println("Time's up!")
			return correctAnswers
		}
	}
}

func main() {
	f, err := os.Open("problem.csv")
	if err != nil {
		log.Fatal(err)
	}

	quiz := NewQuizData()
	quiz.quizReader(csv.NewReader(f))

	correctAnswers := askQuestions(quiz)
	fmt.Printf("You got %d correct answers from %d\n", correctAnswers, len(quiz.Question))
}
