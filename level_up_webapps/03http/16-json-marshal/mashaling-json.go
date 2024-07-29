package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	Name string
	Age  int
}

func main() {

	u := user{
		Name: "Gopher",
		Age:  25,
	}

	// convert user struct to json
	// serialize user struct to json
	data, err := json.Marshal(u)
	if err != nil {
		return
	}

	fmt.Println(string(data))

}
