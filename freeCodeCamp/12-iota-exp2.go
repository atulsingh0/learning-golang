package main

import "fmt"

const (
	isAdmin = 1 << iota // binary shift
	isHeadQuaters
	canSeeFinancials

	canSeeAfrica
	canSeeAsia
	canSeeEurope
	canSeeNA
	canSeeSA
)

func main() {

	fmt.Printf("isAdmin: %b\n", isAdmin)
	fmt.Printf("isHeadQuaters: %b\n", isHeadQuaters)
	fmt.Printf("canSeeFinancials: %b\n", canSeeFinancials)
	fmt.Printf("canSeeAfrica: %b\n", canSeeAfrica)
	fmt.Printf("canSeeAsia: %b\n", canSeeAsia)
	fmt.Printf("canSeeEurope: %b\n", canSeeEurope)

	yourRole := isAdmin | canSeeFinancials | canSeeAsia
	fmt.Printf("\nYour Role is: %b\n", yourRole)

	// check if I can see Europe data
	fmt.Printf("Can see Europe: %v\n", (yourRole&canSeeEurope == canSeeEurope))

	// check if I can see Finantial data
	fmt.Printf("Can see Finantial data: %v\n", (yourRole&canSeeFinancials == canSeeFinancials))

}
