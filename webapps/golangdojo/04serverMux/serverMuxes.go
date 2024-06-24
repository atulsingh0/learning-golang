package main

import (
	"log"
	"net/http"
	"time"
)

type customHandle struct {
	format string
}

// any data type can be handler if it implements ServeHTTP method
func (ch customHandle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tm := time.Now().Format(ch.format)
	w.Write([]byte("The time is: " + tm))
}

func main() {

	mux := http.NewServeMux()
	ch := customHandle{format: time.RFC1123}
	mux.Handle("/", ch)

	log.Print("Listening...")
	http.ListenAndServe(":8001", mux)
}
