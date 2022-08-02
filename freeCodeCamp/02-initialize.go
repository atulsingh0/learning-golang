package main

import "fmt"

// global var
var zed int = 51

var (
	name    string = "atul"
	job     string = "SE"
	address        = "260 Queen St, Toronot"
)

func main() {

	var alpha string
	var beta int = 42
	gamma := 3.14
	delta := "this is me"

	fmt.Printf("%v : %T\n", alpha, alpha)
	fmt.Printf("%v : %T\n", beta, beta)
	fmt.Printf("%v : %T\n", gamma, gamma)
	fmt.Printf("%v : %T\n", delta, delta)

	fmt.Printf("%v : %T\n", zed, zed)

	fmt.Printf("%v : %T\n", name, name)
	fmt.Printf("%v : %T\n", job, job)

	fmt.Printf("%v : %T\n", address, address) // gobal val

	address := "260 King St, Toronto" // local var take predence (called shadowing)
	fmt.Printf("%v : %T\n", address, address)

}
