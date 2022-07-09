package urlhandler

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMapHandler(t *testing.T) {
	// TODO: implement this
	server := InitMux()
	t.Run("load HandlerFunc without errors", func(t *testing.T) {
		pathsToUrls := map[string]string{
			"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
			"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
		}

		map_server := MapHandler(pathsToUrls, server)

		request, _ := http.NewRequest(http.MethodGet, "/yaml-godoc", nil)
		response := httptest.NewRecorder()

		map_server.ServeHTTP(response, request)

		if response.Result().StatusCode != 200 {
			t.Errorf("expect %q, got %q", response.Result().StatusCode, 200)
		}
	})
}

// func TestYAMLHandler(t *testing.T) {
// 	// TODO: implement this
// }
