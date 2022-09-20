package main

import (
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello world!"))
	})

	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		panic(err.Error())
	}

}
