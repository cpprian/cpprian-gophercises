package quizgame

import (
	"encoding/csv"
	"fmt"
	"log"
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

func (q *QuizData) putNewQuestion(line []string) error {
	question, answer := line[0], line[1]
	if len(question) == 0 || len(answer) == 0 {
		return fmt.Errorf("wrong text length")
	}
	q.Question[question] = answer
	return nil
}
