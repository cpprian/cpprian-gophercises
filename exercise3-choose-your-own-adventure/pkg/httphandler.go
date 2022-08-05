package adventure

import (
	"net/http"
	"strings"
)

func (ah *AdventureHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	arg := strings.TrimPrefix(r.URL.Path, "/")

	_, ok := ah.Content[arg]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Unable to find this argument!"))
	}

	w.WriteHeader(http.StatusAccepted)
}
