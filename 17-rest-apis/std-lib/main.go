package main

import (
	"net/http"

)

type homeHandler struct{}

func (h *homeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is my home page"))
}

func main() {
	mux := http.NewServeMux()

	mux.Handle("/", &homeHandler{})
	mux.Handle("/recipes", &RecipesHandler{})
	mux.Handle("/recipes/", &RecipesHandler{})

	http.ListenAndServe(":8080", mux)
}
