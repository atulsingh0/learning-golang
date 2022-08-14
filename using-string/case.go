package main

import (
	"fmt"
	"strings"
)

func main() {

	sampleText := "This is me learning GOlang"

	fmt.Println(strings.Title(sampleText))
	fmt.Println(strings.ToLower(sampleText))
	fmt.Println(strings.ToUpper(sampleText))

	fmt.Println(properTitle(sampleText))

}

func properTitle(str string) string {
	words := strings.Split(str, " ")

	smallLetter := " a an the to of is are am has had have "

	for i, word := range words {
		if !strings.Contains(smallLetter, " "+word+" ") {
			words[i] = strings.Title(word)
		} else {
			words[i] = word
		}
	}

	return strings.Join(words, " ")
}
