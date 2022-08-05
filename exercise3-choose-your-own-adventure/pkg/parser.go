package adventure

import (
	"encoding/json"
	"os"
)

type JsonParser interface {
	Parse(data []byte) error
}

func (adh AdventureHandler) Parse(data []byte) error {
	err := json.Unmarshal(data, &adh.Content)
	return err
}

func InitAdventureHandler() *AdventureHandler {
	return &AdventureHandler{
		AdventureArray{},
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
