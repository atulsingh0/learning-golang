package main

import (
	"encoding/json"
	"fmt"
)

type Article struct {
	Name string `json:"name"`
	// field will not be represent if empty
	Author string `json:"writer,omitempty"`
	// field will not be represent
	Commission float64 `json:"-"`
	draft      bool
}

func main() {

	article := Article{
		Name:       "Json in GO",
		Author:     "Curtis",
		Commission: 12.02,
		draft:      true,
	}

	// Indented JSON print
	data, _ := json.MarshalIndent(article, "", "  ")
	fmt.Println(string(data))

	article = Article{
		Name:       "Json in GO 2",
		Author:     "",
		Commission: 12.02,
		draft:      true,
	}

	// Indented JSON print
	data, _ = json.MarshalIndent(article, "", "  ")
	fmt.Println(string(data))

}
