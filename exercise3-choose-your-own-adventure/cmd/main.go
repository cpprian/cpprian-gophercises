package main

import (
	"log"
	"net/http"

	mypkg "github.com/cpprian/cpprian-gophercises/exercise3-choose-your-own-adventure/pkg"
)

func main() {
	adh, err := mypkg.LoadJsonContent("../gopher.json")
	if err != nil {
		log.Println(err)
	}

	log.Println(http.ListenAndServe(":8080", adh))
}