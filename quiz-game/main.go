package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

type QuestionAnswer struct {
	question string
	answer   string
}

type QuizGame struct {
	qa            []*QuestionAnswer
	correctAnswer int
}

func main() {
	// load user's flags
	var filenameFlag = flag.String("f", "problem.csv", "provide a csv file from which you want to learn")

	flag.Parse()

	// load a csv file
	file, err := os.Open(*filenameFlag)
	if err != nil {
		log.Fatalf("cannot open file: %p", err)
		return
	}
	defer file.Close()

	quizgame := QuizGame{}
	err = quizgame.newDeck(file)
	if err != nil {
		return
	}

	quizgame.startGame()

	fmt.Printf("Your score: %d/%d\n", quizgame.correctAnswer, len(quizgame.qa))
}

func (q *QuizGame) newDeck(f *os.File) error {
	reader := csv.NewReader(f)
	data, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("cannot read csv file: %p", err)
		return err
	}

	for _, row := range data {
		q.qa = append(q.qa, &QuestionAnswer{
			strings.Trim(row[0], " \n"),
			strings.Trim(row[1], " \n"),
		})
	}

	return nil
}

func (q *QuizGame) startGame() {
	reader := bufio.NewReader(os.Stdin)

	for _, card := range q.qa {
		fmt.Printf("%s -> ", card.question)
		text, _ := reader.ReadString('\n')

		if strings.Trim(text, " \n") == card.answer {
			q.correctAnswer += 1
		}
	}
}
