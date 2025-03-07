package main

import (
	"fmt"
	"net/http"
)

func main() {

	// creating custom server
	mux := http.NewServeMux()

	// path type of url
	mux.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "List of Users")

	})

	// subtree type of url
	mux.HandleFunc("/product/", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, fmt.Sprintf("%s page not found", r.URL.String()), http.StatusNotFound) // 404

	})

	mux.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/users", http.StatusMovedPermanently) // 301
	})

	http.ListenAndServe(":3000", mux)
}
