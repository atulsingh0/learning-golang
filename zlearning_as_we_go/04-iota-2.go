package main

import "fmt"

const (
	i = iota << 1 //
	j             //
	k
)

func main() {

	fmt.Println(i, j, k)

}
