package main

import (
	"fmt"
	"net/http"

	mypkg "github.com/cpprian/cpprian-gophercises/exercise4-htmlparser/pkg"
)

func main() {
	href := mypkg.NewParser()

	f, err := http.Get("https://www.calhoun.io")
	if err != nil {
		error.Error(err)
	}

	href.Parse(f.Body)
	fmt.Println(href)
}
