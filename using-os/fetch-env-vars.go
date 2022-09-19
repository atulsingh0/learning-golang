package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	// fetch all the env variable
	env := os.Environ()

	// loop over env var
	for _, v := range env {

		// split the key and value
		data := strings.Split(v, "=")
		fmt.Println(data[0], "--->", data[1])
	}
}
