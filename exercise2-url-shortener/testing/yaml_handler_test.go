package urlhandler_testing

import (
	"net/http"
	"net/http/httptest"
	"testing"

	handle "github.com/cpprian/cpprian-gophercises/exercise2-url-shortener/handle"
)

func TestYAMLHandler(t *testing.T) {

	server := handle.InitMux()
	t.Run("create HandlerFunc from YAML file", func(t *testing.T) {
		map_server, err := handle.YAMLHandler([]byte("../store/yaml_test.yml"), server)
		if err != nil {
			t.Error(err)
		}

		request, _ := http.NewRequest(http.MethodGet, "/lofi-girl", nil)
		response := httptest.NewRecorder()
		map_server.ServeHTTP(response, request)

		assertResponseCode(t, response.Result().StatusCode, 301)
	})
}
