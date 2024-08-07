package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

/*
Simply run "go get pkg.name" to download the package.
and added to your go.mod file.
Remember to tun the command "go mod tidy" to update the go.mod file.

Running "go mod vendor" will download the package to the vendor folder.
*/

func main() {

	srv := mux.NewRouter()

	srv.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, world!")
	})

	http.ListenAndServe(":3000", srv)
}
