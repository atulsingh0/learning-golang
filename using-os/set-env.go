package main

import (
	"fmt"
	"os"
)

func main() {

	// Setting env var "ABCD"
	os.Setenv("ABCD", "123")
	fmt.Println(os.Getenv("ABCD"))
}
