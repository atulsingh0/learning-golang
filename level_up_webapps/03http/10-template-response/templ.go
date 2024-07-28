package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {

	tmpl, err := template.New("test").Parse("<h1>Hello {{.}}!</h2>")
	if err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err = tmpl.Execute(w, `Gophers`)
		if err != nil {
			return
		}
	})

	http.ListenAndServe(":3000", mux)

}
