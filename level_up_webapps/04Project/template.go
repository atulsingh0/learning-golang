package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var templates = template.Must(template.New("t").ParseGlob("templates/**/*.html"))
var errorTemplate = `<html><body><h1>Template Error: %s</h1?</body><p>Error: %s</p></body></html>`

func renderTemplate(w http.ResponseWriter, r *http.Request, name string, data interface{}) {
	err := templates.ExecuteTemplate(w, name, data)
	if err != nil {
		http.Error(w, fmt.Sprintf(errorTemplate, name, err), http.StatusInternalServerError)
	}
}
