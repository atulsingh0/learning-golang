package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Person struct {
	Name      string `json:"name"`
	Height    string `json:"height"`
	Mass      string `json:"mass"`
	EyeColor  string `json:"eye_color"`
	BirthYear string `json:"birth_year"`
	Gender    string `json:"gender"`
}

func main() {
	p := unmarshal(2)
	fmt.Printf("%+v\n", p)

	p = helperUnmarshal(-1)
	fmt.Printf("%+v\n", p)
}

// function to call the API and unmarshal the response
// error handling are done in the function itself
func unmarshal(id int) Person {
	r, err := http.Get(fmt.Sprintf("https://swapi.dev/api/people/%d", id))
	defer r.Body.Close()

	if r.StatusCode != 200 {
		log.Printf("Error: (%d) %s", r.StatusCode, r.Status)
	}

	data, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error: %s", err)
	}

	p := Person{}
	err = json.Unmarshal(data, &p)
	if err != nil {
		log.Printf("Error: %s", err)
	}
	return p
}

func helperUnmarshal(id int) Person {
	r, err := http.Get(fmt.Sprintf("https://swapi.dev/api/people/%d", id))
	defer r.Body.Close()

	if r.StatusCode != 200 {
		log.Printf("Error: (%d) %s", r.StatusCode, r.Status)
	}

	data, err := io.ReadAll(r.Body)
	errorParser("reading json response:", err)

	p := Person{}
	err = json.Unmarshal(data, &p)
	errorParser("Unmarshal json response:", err)

	return p
}

func errorParser(msg string, err error) {
	if err != nil {
		log.Printf("!! Error: %s : %s", msg, err.Error())
	}
}
