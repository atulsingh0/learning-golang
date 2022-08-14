package main

import (
	"fmt"
	"strings"
)

func main() {

	sampleTxt := "     this is our    test           "

	fmt.Printf("%q\n", strings.Trim(sampleTxt, " "))
	fmt.Printf("%q\n", strings.TrimLeft(sampleTxt, " "))
	fmt.Printf("%q\n", strings.TrimRight(sampleTxt, " "))

	fmt.Printf("%q\n", strings.TrimPrefix(strings.Trim(sampleTxt, " "), "this"))

}
