package main

import (
	"fmt"
	"net/http"
)

type user struct {
	Name string
	Age  int
}

func (u user) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome!, ", u.Name)
}

func NewUser(name string, age int) *user {
	return &user{Name: name, Age: age}
}

func main() {
	// creating custom server
	mux := http.NewServeMux()

	u := NewUser("Gopher", 25)
	mux.Handle("/", u)

	// path type of url
	mux.HandleFunc("/data", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "data list")

	})

	http.ListenAndServe(":3000", mux)

}
