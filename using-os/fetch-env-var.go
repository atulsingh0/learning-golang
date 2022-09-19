// Go Playground
// https://go.dev/play/p/9wqZwtlDdQk

package main

import (
	"fmt"
	"os"
)

func main() {

	// fetch Language env var
	// If env exists, return value else return empty
	fmt.Println("PAGER env:", os.Getenv("PAGER"))

	// to check if Env exists or not
	// if env var does not exists, retrun false else true
	if val, ok := os.LookupEnv("PAGER"); ok == false {
		fmt.Println("Env var does not exists.")
	} else {
		fmt.Println("Env val exists and value is:", val)
	}

}
