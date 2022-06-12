package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

type QuizData struct {
	Question map[string]string
}

var (
	filename  string
	timelimit int
)

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

func askQuestions(q *QuizData, limit int) int {
	correctAnswers := 0
	index := 1
	ch := make(chan bool)
	timer := time.NewTimer(time.Duration(limit) * time.Second)

	go func() {
		for question, answer := range q.Question {
			var a string
			fmt.Printf("%d. %s\n", index, question)
			fmt.Scanln(&a)
			if a == answer {
				ch <- true
			} else {
				ch <- false
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
	flag.IntVar(&timelimit, "time", 30, "use for time limit for all of your questions to solve (by default 30)")
	flag.StringVar(&filename, "filename", "problem.csv", "choose which file you want to use to play a quiz")
	flag.Parse()

	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	quiz := NewQuizData()
	quiz.quizReader(csv.NewReader(f))

	correctAnswers := askQuestions(quiz, timelimit)
	fmt.Printf("You got %d correct answers from %d\n", correctAnswers, len(quiz.Question))
}
