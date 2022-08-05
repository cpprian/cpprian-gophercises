package adventure

import (
	"encoding/json"
	"net/http"
	"os"
)

type AdventureStruct struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []struct {
		Text string `json:"text"`
		Arc  string `json:"arc"`
	} `json:"options"`
}

type AdventureArray map[string]AdventureStruct

type JsonParser interface {
	Parse(data []byte) error
}

type AdventureHandler struct {
	Hello AdventureArray
	http.Handler
}

func (adh AdventureHandler) Parse(data []byte) error {
	err := json.Unmarshal(data, &adh.Hello)
	return err
}

func InitAdventureHandler() *AdventureHandler {
	return &AdventureHandler{
		AdventureArray{},
		nil,
	}
}

func LoadJsonContent(filename string) (*AdventureHandler, error) {
	newAdventure := InitAdventureHandler()

	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	err = newAdventure.Parse(data)
	return newAdventure, err
}
