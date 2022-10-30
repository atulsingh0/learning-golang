package main

import (
	"customdatatypes/org"
	"os"

	"fmt"
)

func main() {
	fmt.Println("Hi")
	p := org.NewPerson("Priya", "Chaudhary")

	fmt.Println("ID:", p.ID())
	fmt.Println(p.FullName())

	// setting twitter
	err := p.SetTwitter(org.TwitterHandler("@priya"))
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println(string(p.GetTwitter()))       // returns TwitterHandler type
	fmt.Println(p.GetTwitter().RedirectUrl()) // Have access on all the methods exposed with TwitterHandler type

	// fmt.Println(p.TwitterHandler)

}
