package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"log"
	"net/http"

	mypkg "github.com/cpprian/cpprian-gophercises/exercise4-htmlparser/pkg"
	builder "github.com/cpprian/cpprian-gophercises/exercise5-sitemap-builder/pkg"
)

func main() {
	flag.IntVar(&builder.Maxdepth, "d", 3, "how many times the program can come in deeper in links")
	flag.StringVar(&builder.Website, "w", "https://www.calhoun.io", "choose your website to create xml")
	flag.Parse()

	fmt.Println(builder.Maxdepth)

	resp, err := http.Get(builder.Website)
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
	b.Build(*r, builder.Website, 0, builder.Maxdepth)

	hello, err := xml.MarshalIndent(b, " ", "   ")
	if err != nil {
		log.Printf("Marshal error: %v", err)
	}

	fmt.Print(xml.Header)
	fmt.Println(string(hello))
}
