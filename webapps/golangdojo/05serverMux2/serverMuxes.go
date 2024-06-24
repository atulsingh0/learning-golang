package main

import (
	"log"
	"net/http"
	"time"
)

func handleRoot(w http.ResponseWriter, r *http.Request) {
	tm := time.Now().Format(time.RFC1123)
	log.Println("handling root...")
	w.Write([]byte("The time is: " + tm))
}

func main() {

	mux := http.NewServeMux()
	// Functions as handlers
	mux.HandleFunc("/", handleRoot)

	log.Print("Listening on port 8001...")
	http.ListenAndServe(":8001", mux)
}
