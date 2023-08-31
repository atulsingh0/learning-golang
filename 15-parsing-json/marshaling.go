package main

import (
	"encoding/json"
	"fmt"
)

type Article struct {
	Name   string
	Author string
	draft  bool
}

func main() {

	article := Article{
		Name:   "Json in GO",
		Author: "Curtis",
		draft:  true,
	}

	// private keys are not exporeted which is "draft" in our case
	data, err := json.Marshal(article)
	if err != nil {
		fmt.Println("Unable to marshal", err.Error())
	} else {
		fmt.Println(string(data))
	}

	// Indented JSON print
	data, _ = json.MarshalIndent(article, "", "  ")
	fmt.Println(string(data))

}
