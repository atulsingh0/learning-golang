package main

import (
	"log"
	"net/http"
	"reflect"
	"runtime"
	"time"
)

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/", hello)

	// using a middleware
	mux.HandleFunc("/hello", logg(hello))

	log.Fatal(http.ListenAndServe(":8080", mux))
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

func logg(h http.HandlerFunc) http.HandlerFunc {
	// this is a closure
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Request: ", r.URL.Path)
		start := time.Now()
		h(w, r)
		name := runtime.FuncForPC(reflect.ValueOf(h).Pointer()).Name()
		log.Printf("%s took %s", name, time.Since(start))
	}
}
