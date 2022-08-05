package adventure_test

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	mypkg "github.com/cpprian/cpprian-gophercises/exercise3-choose-your-own-adventure/pkg"
)

func TestAdventureHandlerServeHTTP(t *testing.T) {
	log.Println("Running TestAdventureHandlerServeHTTP...")

	t.Run("test request with existing key from json", func(t *testing.T) {
		adl := AdventureHandlerLoader(t)

		request, _ := http.NewRequest(http.MethodGet, "/new-york", nil)
		response := httptest.NewRecorder()

		adl.ServeHTTP(response, request)

		got := response.Result().StatusCode
		want := http.StatusAccepted

		AssertStatusRequest(t, got, want)
	})

	t.Run("test request with non-existing key", func(t *testing.T) {
		adl := AdventureHandlerLoader(t)

		request, _ := http.NewRequest(http.MethodGet, "/arizona", nil)
		response := httptest.NewRecorder()

		adl.ServeHTTP(response, request)

		got := response.Result().StatusCode
		want := http.StatusBadRequest

		AssertStatusRequest(t, got, want)
	})
}

func AssertStatusRequest(t testing.TB, got, want int) {
	if got != want {
		t.Errorf("want '%v'\n\n got '%v'\n", want, got)
	}
}

func AdventureHandlerLoader(t testing.TB) *mypkg.AdventureHandler {
	adl, err := mypkg.LoadJsonContent("../gopher.json")
	if err != nil {
		t.Errorf("%v\n", err)
	}
	return adl
}