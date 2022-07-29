package urlhandler_testing

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	handle "github.com/cpprian/cpprian-gophercises/exercise2-url-shortener/handle"
)

func TestMapHandler(t *testing.T) {
	
	server := handle.InitMux()
	t.Run("load HandlerFunc", func(t *testing.T) {
		pathsToUrls := map[string]string{
			"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
			"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
		}

		map_server := handle.MapHandler(pathsToUrls, server)

		request, _ := http.NewRequest(http.MethodGet, "/yaml-godoc", nil)
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