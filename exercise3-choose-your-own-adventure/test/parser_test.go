package adventure_test

import (
	"testing"

	mypkg "github.com/cpprian/cpprian-gophercises/exercise3-choose-your-own-adventure/pkg"
)

func TestLoadJsonContent(t *testing.T) {

	t.Run("convert simple json text to struct", func(t *testing.T) {
		testedAdventureHandler, err := mypkg.LoadJsonContent("../gopher.json")
		if err != nil {
			t.Errorf("Can't load json file -> %v", err)
		}

		if len(testedAdventureHandler.Hello) != 7 {
			t.Errorf("Want %d, got %d", 7, len(testedAdventureHandler.Hello))
		}

		if _, mp := testedAdventureHandler.Hello["new-york"]; mp == false {
			t.Errorf("Element new-york doesn't exist")
		}
	})
}
