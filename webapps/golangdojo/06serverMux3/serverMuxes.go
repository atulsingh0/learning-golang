package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// returning http.Handler
func timeHandler(format string) http.Handler {
	log.Println("Registering root...")

	fn := func(w http.ResponseWriter, r *http.Request) {
		tm := time.Now().Format(format)
		w.Write([]byte("The time is: " + tm))
	}
	return http.HandlerFunc(fn)
}

// Using implicit conversion
// which simplifies the code block
func userHandler(user string) http.HandlerFunc {
	log.Println("Registering /user...")

	return func(w http.ResponseWriter, r *http.Request) {
		tm := time.Now()
		fmt.Fprintf(w, "User %s is logged in from: %s\n", user, tm)
	}

}

func main() {

	mux := http.NewServeMux()

	// Functions as handlers which accept inputs
	mux.Handle("/", timeHandler(time.RFC1123))
	mux.Handle("/user", userHandler("Atul"))

	log.Print("Listening on port 8001...")
	http.ListenAndServe(":8001", mux)
}
