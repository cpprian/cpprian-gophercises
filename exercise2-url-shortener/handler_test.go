package urlhandler

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestMapHandler(t *testing.T) {
	
	server := InitMux()
	t.Run("load HandlerFunc", func(t *testing.T) {
		pathsToUrls := map[string]string{
			"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
			"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
		}

		map_server := MapHandler(pathsToUrls, server)

		request, _ := http.NewRequest(http.MethodGet, "/yaml-godoc", nil)
		response := httptest.NewRecorder()

		map_server.ServeHTTP(response, request)
		assertResponseCode(t, response.Result().StatusCode, 301)
	})
}

func TestYAMLtoMap(t *testing.T) {

	t.Run("give a string with YAML syntax", func(t *testing.T) {
		yml := `
- path: /urlshort
  url: https://github.com/gophercises/urlshort
- path: /urlshort-final
  url: https://github.com/gophercises/urlshort/tree/solution
`
		want := map[string]string{
			"/urlshort": "https://github.com/gophercises/urlshort",
			"/urlshort-final": "https://github.com/gophercises/urlshort/tree/solution",
		}

		got, err := YAMLtoMap([]byte(yml)) 
		if err != nil {
			t.Error(err)
		}

		assertMapPaths(t, got, want)
	})

	t.Run("provide an YAML file", func(t *testing.T) {
		want := map[string]string {
			"/go": "https://go.dev",
			"/urlshortener": "https://github.com/gophercises/urlshort",
			"/lofi-girl": "https://lofimusic.app/lofigirl",
		}

		got, err := YAMLtoMap([]byte("yaml_test.yml")) 
		if err != nil {
			t.Error(err)
		} 

		assertMapPaths(t, got, want)
	})
}

func TestYAMLHandler(t *testing.T) {

	server := InitMux()
	t.Run("create HandlerFunc from YAML file", func(t *testing.T) {
		map_server, err := YAMLHandler([]byte("yaml_test.yml"), server)
		if err != nil {
			t.Error(err)
		} 

		request, _ := http.NewRequest(http.MethodGet, "/lofi-girl", nil)
		response := httptest.NewRecorder()
		map_server.ServeHTTP(response, request)

		assertResponseCode(t, response.Result().StatusCode, 301)
	})	
}

func assertResponseCode(t testing.TB, got, want int) {
	t.Helper()
	if got != 301 {
		t.Errorf("expect %d, got %d", want, got)
	}
}

func assertMapPaths(t testing.TB, got, want map[string]string) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("want %v, got %v", want, got)
	}
}
