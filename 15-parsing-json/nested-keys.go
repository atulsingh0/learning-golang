package main

import (
	"encoding/json"
	"fmt"
)

type Author struct {
	Name    string
	Country string
}

type Article struct {
	Name   string
	Author Author
	draft  bool
}

func main() {

	article := Article{
		Name: "Json in GO",
		Author: Author{
			Name:    "Curtis",
			Country: "US",
		},
		draft: true,
	}

	// Indented JSON print
	data, _ := json.MarshalIndent(article, "", "  ")
	fmt.Println(string(data))

}
