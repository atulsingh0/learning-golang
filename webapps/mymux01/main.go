package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {

	// define a custom router
	srv := &http.Server{
		Addr:         ":8001",
		Handler:      nil,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}
	s := http.NewServeMux()
	srv.Handler = s

	s.HandleFunc("/", handleFunction)

	s.HandleFunc("/pic", parserPic)
	err := srv.ListenAndServe()
	if err != nil {
		log.Printf("server startup error: %s", err.Error())
		return
	}
}

func parserPic(w http.ResponseWriter, r *http.Request) {
	log.Printf("Call processed by parserPic: %s\n", r.Method)
	fmt.Fprint(w, getRandomPic())
}

func handleFunction(w http.ResponseWriter, r *http.Request) {
	log.Printf("Call processed by handleFunction: %s\n", r.Method)
	fmt.Fprint(w, "Welcome to Hello World")
}
