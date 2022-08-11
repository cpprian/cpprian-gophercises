package main

import (
	"fmt"
	"log"
	"net/http"

	mypkg "github.com/cpprian/cpprian-gophercises/exercise4-htmlparser/pkg"
)

func main() {
	url := "https://go.dev"
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("GET error: %v", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("Status error: %v", resp.StatusCode)
		return
	}

	r := mypkg.NewParser()	
	r.Parse(resp.Body)

	for i := 0; i < len(*r); i++ {
		fmt.Println((*r)[i].Href, (*r)[i].Text)
	}
}
