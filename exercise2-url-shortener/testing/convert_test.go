package urlhandler_testing

import (
	"testing"

	handle "github.com/cpprian/cpprian-gophercises/exercise2-url-shortener/handle"
)

func TestConverToMapPaths(t *testing.T) {

	t.Run("give a string with YAML syntax", func(t *testing.T) {
		yml := `
- path: /urlshort
  url: https://github.com/gophercises/urlshort
- path: /urlshort-final
  url: https://github.com/gophercises/urlshort/tree/solution
`
		want := map[string]string{
			"/urlshort":       "https://github.com/gophercises/urlshort",
			"/urlshort-final": "https://github.com/gophercises/urlshort/tree/solution",
		}

		yaml_struct := &handle.YamlList{}
		got, err := yaml_struct.ConvertToMapPaths([]byte(yml))
		if err != nil {
			t.Error(err)
		}

		assertMapPaths(t, got, want)
	})

	t.Run("provide an YAML file", func(t *testing.T) {
		want := map[string]string{
			"/go":           "https://go.dev",
			"/urlshortener": "https://github.com/gophercises/urlshort",
			"/lofi-girl":    "https://lofimusic.app/lofigirl",
		}

		yaml_struct := &handle.YamlList{}
		got, err := yaml_struct.ConvertToMapPaths([]byte("../store/yaml_test.yml"))
		if err != nil {
			t.Error(err)
		}

		assertMapPaths(t, got, want)
	})

	t.Run("give a string with JSON syntax", func(t *testing.T) {
		jsn := `
[
	{
		"path": "/go",
		"url": "https://go.dev"
	},
	{
		"path": "/urlshortener",
		"url": "https://github.com/gophercises/urlshort"
	},
	{
		"path": "/lofi-girl",
		"url": "https://lofimusic.app/lofigirl"
	}
]	
`
		want := map[string]string{
			"/go":           "https://go.dev",
			"/urlshortener": "https://github.com/gophercises/urlshort",
			"/lofi-girl":    "https://lofimusic.app/lofigirl",
		}

		json_struct := &handle.JsonList{}
		got, err := json_struct.ConvertToMapPaths([]byte(jsn))
		if err != nil {
			t.Error(err)
		}

		assertMapPaths(t, got, want)
	})

	t.Run("provide a JSON file", func(t *testing.T) {
		want := map[string]string{
			"/go":           "https://go.dev",
			"/urlshortener": "https://github.com/gophercises/urlshort",
			"/lofi-girl":    "https://lofimusic.app/lofigirl",
		}

		json_struct := &handle.JsonList{}
		got, err := json_struct.ConvertToMapPaths([]byte("../store/json_test.json"))
		if err != nil {
			t.Error(err)
		}

		assertMapPaths(t, got, want)
	})
}
