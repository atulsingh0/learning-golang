package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Name  string
	Age   int
	Class int
	Marks []struct {
		Subject string
		Mark    int
	}
}

func main() {

	st := Student{
		Name:  "Gopher",
		Age:   25,
		Class: 12,
		Marks: []struct {
			Subject string
			Mark    int
		}{
			{"Math", 95},
			{"Science", 92},
			{"English", 91},
		},
	}

	// convert user struct to json
	// serialize user struct to json
	data, err := json.MarshalIndent(st, "", "  ")
	if err != nil {
		return
	}

	fmt.Println(string(data))

}
