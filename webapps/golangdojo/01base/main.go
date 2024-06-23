package main

import (
	"log"
	"net/http"
)

func HelloWorldPage(w http.ResponseWriter, r *http.Request) {
	log.Println("Called to /")
	_, _ = w.Write([]byte("Hi, You have reached to hello world. \t"))
	//_, _ = fmt.Fprint(w, "This is a second line")
}

func main() {

	// HandleFunc ise DefailtServerMux to build the server
	http.HandleFunc("/", HelloWorldPage)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal("Server side issue:", err.Error())
	}

}
