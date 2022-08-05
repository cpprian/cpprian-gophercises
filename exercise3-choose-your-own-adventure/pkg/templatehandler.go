package adventure

import (
	"html/template"
	"net/http"
)

func ExecuteTemplates(w http.ResponseWriter, as AdventureStruct) {
	tmpl, err := template.ParseGlob("../templates/*.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	tmpl.ExecuteTemplate(w, "adventure", as)
}
