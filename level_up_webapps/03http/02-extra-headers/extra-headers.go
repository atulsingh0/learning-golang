package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.Header().Set("X-Content-Type-Options", "nosniff")

		// Add the keys and values to the header.
		w.Header().Add("X-Frame-Options", "deny")
		w.Header().Add("X-Frame-Options", "sameorigin")

		// Set the key and value to the header, overriding any existing values.
		w.Header().Set("X-Frame-Options", "deny")

		fmt.Fprint(w, "Hello Gophers")
	})

	http.ListenAndServe(":8080", nil)
}
