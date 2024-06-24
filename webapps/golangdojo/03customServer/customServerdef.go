package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {

	srv := &http.Server{
		Addr:         ":8001",
		Handler:      nil,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	http.HandleFunc("/", handleFunction)
	srv.ListenAndServe()
}

func handleFunction(w http.ResponseWriter, r *http.Request) {
	log.Printf("Call processed by handleFunction: %s\n", r.Method)
	switch r.URL.Path {
	case "/":
		fmt.Fprint(w, "Welcome to Hello World")
	case "/user":
		log.Println("Sleeping for 2 second intentionally.")
		time.Sleep(2 * time.Second)
		fmt.Fprint(w, "This will never be reached")
	default:
		fmt.Fprint(w, "Requested URI is not found.")
	}

}
