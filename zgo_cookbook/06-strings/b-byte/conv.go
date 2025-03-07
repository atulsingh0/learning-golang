package main

import "fmt"

func main() {
	line := "Hello, world! How are you today?"

	bytes := []byte(line)
	fmt.Println(bytes)

	str := string(bytes)
	fmt.Println(str)
}
