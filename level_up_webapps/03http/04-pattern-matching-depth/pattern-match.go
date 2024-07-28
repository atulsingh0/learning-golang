package main

import (
	"fmt"
	"net/http"
)

func main() {

	// creating custom server
	mux := http.NewServeMux()

	// in case of no path, it will handle all the requests
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello Gophers")
	})

	// path type of url
	mux.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "List of Users")

	})

	// subtree type of url
	mux.HandleFunc("/product/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "List of Products")

	})

	http.ListenAndServe(":3000", mux)
}
