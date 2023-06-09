package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {

	fmt.Println(time.Now())
	fmt.Println(time.Now().Date())
	fmt.Println(time.Now().Year())

	sentence := "asta la vista"

	fmt.Println(strings.Replace(sentence, "a", "Z", -1)) // replace all occurance
	fmt.Println(strings.Replace(sentence, "a", "P", 2))  // replace 2 occurance

}
