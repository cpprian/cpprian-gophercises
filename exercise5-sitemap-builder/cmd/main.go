package main

import (
	"encoding/xml"
	"fmt"
	"log"
	"net/http"

	mypkg "github.com/cpprian/cpprian-gophercises/exercise4-htmlparser/pkg"
	builder "github.com/cpprian/cpprian-gophercises/exercise5-sitemap-builder/pkg"
)

func main() {
	url := "https://www.calhoun.io"
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

	b := &builder.XmlContent{}
	b.Build(*r, url)

	hello, err := xml.MarshalIndent(b, " ", " ")
	if err != nil {
		log.Printf("Marshal error: %v", err)
	}

	fmt.Print(xml.Header)
	fmt.Println(string(hello))
}
