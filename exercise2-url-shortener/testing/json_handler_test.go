package urlhandler_testing

import (
	"net/http"
	"net/http/httptest"
	"testing"

	handle "github.com/cpprian/cpprian-gophercises/exercise2-url-shortener/handle"
)

func TestJSONHandler(t *testing.T) {

	server := handle.InitMux()
	t.Run("create HandlerFunc from JSON file", func(t *testing.T) {
		map_server, err := handle.YAMLHandler([]byte("../store/json_test.json"), server)
		if err != nil {
			t.Error(err)
		}

		request, _ := http.NewRequest(http.MethodGet, "/lofi-girl", nil)
		response := httptest.NewRecorder()
		map_server.ServeHTTP(response, request)

		assertResponseCode(t, response.Result().StatusCode, 301)
	})
}
