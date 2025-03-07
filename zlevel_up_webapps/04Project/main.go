package main

import (
	"log"
	"net/http"
)

// main is the entry point for the application.
// Program must be run with this directory as the working directory.
func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, r, "index/home", nil)
	})

	mux.Handle("/assets/",
		http.StripPrefix(
			"/assets/",
			http.FileServer(http.Dir("assets/")),
		),
	)

	err := http.ListenAndServe(":9000", mux)
	if err != nil {
		log.Fatal(err)
	}
}
