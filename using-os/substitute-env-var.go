// Go Playground
// https://go.dev/play/p/k2p8TN1KtYS

package main

import (
	"fmt"
	"os"
)

func main() {
	// setting 2 env
	os.Setenv("NAME", "atul")
	os.Setenv("LOC", "/usr/atul")

	// replace the $ENV with its value
	fmt.Println(os.ExpandEnv("$NAME lives in ${LOC}."))
	fmt.Println(os.ExpandEnv("Home dir is ${HOME}."))
	fmt.Println(os.ExpandEnv("Current dir is ${PWD}."))

}
