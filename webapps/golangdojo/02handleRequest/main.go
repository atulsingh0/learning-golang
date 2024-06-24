package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", handleFunction)
	http.ListenAndServe(":8001", nil)
}

func handleFunction(w http.ResponseWriter, r *http.Request) {
	log.Printf("Call processed by handleFunction: %s\n", r.Method)
	switch r.URL.Path {
	case "/":
		fmt.Fprint(w, "Welcome to Hello World")
	case "/user":
		fmt.Fprint(w, "Just kidding, there is no User")
	default:
		fmt.Fprint(w, "Requested URI is not found.")
	}

}
