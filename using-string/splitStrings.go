package main

import (
	"fmt"
	"strings"
)

func main() {

	ourString := "This is a string!"

	stringsCollection := strings.Split(ourString, " ")

	for i := range stringsCollection {
		fmt.Println(i, stringsCollection[i])
	}

	ourString = "This is me learning golang"
	stringsCollection = strings.SplitN(ourString, " ", 3) // only 3 split

	for i := range stringsCollection {
		fmt.Println(stringsCollection[i])
	}

}
