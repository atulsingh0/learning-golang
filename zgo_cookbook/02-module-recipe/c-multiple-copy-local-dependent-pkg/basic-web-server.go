package main

import (
	"fmt"
	"net/http"

	m174 "github.com/gorilla/mux/174"
	m181 "github.com/gorilla/mux/181"
)

/*
Simply run "go get pkg.name" to download the package.
and added to your go.mod file.
Remember to tun the command "go mod tidy" to update the go.mod file.

Running "go mod vendor" will download the package to the vendor folder.
*/

func main() {

	srv1 := m181.NewRouter()
	srv2 := m174.NewRouter()

	srv1.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, world! from 181")
	})

	srv2.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, world! from 174")
	})

	http.ListenAndServe(":3000", srv1)
	// below line never executed
	// http.ListenAndServe(":3001", srv2)

}
