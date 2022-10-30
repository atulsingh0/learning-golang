package main

import (
	"customdatatypes/org"
	"os"

	"fmt"
)

func main() {
	p := org.NewPerson("Priya", "Chaudhary", org.NewSIN("123-123-455"))

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
	// fmt.Println(p.Citizen.Country())

	// fmt.Println(p.TwitterHandler)

	// --------- Europe
	fmt.Println()

	st := org.NewPerson("Adam", "Denver", org.NewEUSIN("999-999-8888", "germany"))

	fmt.Println("ID:", st.ID())

	fmt.Println(st.FullName())

	// setting twitter
	err = st.SetTwitter(org.TwitterHandler("@adamden"))
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println(string(st.GetTwitter()))       // returns TwitterHandler type
	fmt.Println(st.GetTwitter().RedirectUrl()) // Have access on all the methods exposed with TwitterHandler type
	fmt.Printf("%#v\n", st)
	fmt.Printf("ID: %#v\n", st.ID())

}
