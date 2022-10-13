package main

import "fmt"

type Name struct {
	fName string
	mName string
	lName string
}

type Contact struct {
	phone string
}

type Business struct {
	name    Name
	address string
	contact Contact
	dob     DOB
}

type DOB string

func (b *Business) printInfo() {
	fmt.Printf("Name: %s %s %s\n", b.name.fName, b.name.mName, b.name.lName)
	fmt.Printf("Contact: %s\n", b.contact.phone)
	fmt.Printf("Address: %s\n", b.address)
	fmt.Printf("Date of Birth: %s\n", b.dob)
}

func main() {

	user := Business{
		name: Name{
			fName: "Akash",
			mName: "Akbar",
			lName: "Anthony",
		},
		address: "420, Khao gali, chor bazar",
		contact: Contact{
			phone: "143-420-9211",
		},
		dob: DOB(
			"12-12-1972",
		),
	}

	user.printInfo()

	// Changing the fName, DOB, and contact
	user.name.fName = "Amar"
	user.dob = "11-12-1982"
	user.contact = Contact{
		phone: "111-222-3333",
	}

	// print the user info again
	fmt.Println("\nChanged the First Name, DOB and Contact")
	user.printInfo()

}
