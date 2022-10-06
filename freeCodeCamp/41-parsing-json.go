package main

import (
	"encoding/json"
	"fmt"
)

// Dummy Json file
const (
	text = `{
		"id": 1,
		"title": "iPhone 9",
		"description": "An apple mobile which is nothing like apple",
		"price": 549,
		"discountPercentage": 12.96,
		"rating": 4.69,
		"stock": 94,
		"brand": "Apple",
		"category": "smartphones",
		"thumbnail": "https://dummyjson.com/image/i/products/1/thumbnail.jpg",
		"images": [
		  "https://dummyjson.com/image/i/products/1/1.jpg",
		  "https://dummyjson.com/image/i/products/1/2.jpg",
		  "https://dummyjson.com/image/i/products/1/3.jpg",
		  "https://dummyjson.com/image/i/products/1/4.jpg",
		  "https://dummyjson.com/image/i/products/1/thumbnail.jpg"
		]
	  }`
)

// Creating a struct to read the json

type mobile struct {
	Id     int      `json:"id"` // var name should be exported (start with cap letter) to access it outside
	Title  string   `json:"title"`
	Images []string `json:"images"`
}

func main() {

	data := []byte(text)

	// Initializing the variable mb
	mb := mobile{}

	// Parsing the json
	err := json.Unmarshal(data, &mb)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(mb.Id)
	fmt.Println(mb.Title)
	fmt.Println(mb.Images)

}
