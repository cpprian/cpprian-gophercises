package main

import (
	"bufio"
	"os"

	mypkg "github.com/cpprian/cpprian-gophercises/exercise4-htmlparser/pkg"
)

func main() {
	href := mypkg.NewParser()

	f, err := os.Open("../testing/ex2.html")
	if err != nil {
		error.Error(err)
	}
	read := bufio.NewReader(f)
	href.Parse(read)
}
