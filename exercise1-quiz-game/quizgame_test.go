package quizgame

import (
	"encoding/csv"
	"log"
	"os"
	"reflect"
	"testing"
)

func TestQuizReader(t *testing.T) {

	t.Run("read data from csv file", func(t *testing.T) {
		want := &QuizData{
			Question: map[string]string{
				"5+5": "10",
				"7+3": "10",
				"1+1": "2",
				"8+3": "11",
				"1+2": "3",
				"8+6": "14",
				"3+1": "4",
				"1+4": "5",
				"5+1": "6",
				"2+3": "5",
				"3+3": "6",
				"2+4": "6",
				"5+2": "7",
			},
		}

		f, err := os.Open("problem.csv")
		if err != nil {
			log.Fatal(err)
		}
		got := NewQuizData()
		got.quizReader(csv.NewReader(f))

		AssertEqual(t, got, want)
	})
}

func AssertEqual[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if !reflect.DeepEqual(want, got) {
		t.Errorf("got %v, want %v", got, want)
	}
}
