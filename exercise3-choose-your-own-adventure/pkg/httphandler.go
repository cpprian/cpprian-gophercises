package adventure

import (
	"net/http"
	"reflect"
	"strings"
)

func (ah *AdventureHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	arg := strings.TrimPrefix(r.URL.Path, "/")
	if reflect.DeepEqual(arg, "") {
		arg = "intro"
	}

	val, ok := ah.Content[arg]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Unable to find this argument!"))
		return
	}

	w.WriteHeader(http.StatusAccepted)
	ExecuteTemplates(w, val)
}
