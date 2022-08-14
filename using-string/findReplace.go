package main

import (
	"fmt"
	"strings"
)

func main() {

	sampleString := "I am learning golang after spending lot of time working with python."

	searchItem := "go"

	res := strings.Contains(sampleString, searchItem)
	fmt.Printf("String containes \"%v\": %v\n", searchItem, res)

	res = strings.HasPrefix(sampleString, searchItem)
	fmt.Printf("String Starts with \"%v\": %v\n", searchItem, res)

	// res = strings.Replace(sampleString, searchItem, "this", 1)
	replaceWith := "me"
	result := strings.ReplaceAll(sampleString, "go", replaceWith)

	fmt.Println(result)

}
