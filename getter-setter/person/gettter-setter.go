package person

import "fmt"

type name struct {
	fName string
	mName string
	lName string
}

type contact struct {
	phone string
}

type Business struct {
	name    name
	address string
	contact contact
	dob     dob
}

type dob string

// setter functions
func (b *Business) SetName(fName string, mName string, lName string) {
	b.name.fName = fName
	b.name.mName = mName
	b.name.lName = lName
}

func (b *Business) SetDOB(dtOfBirth dob) {
	b.dob = dtOfBirth
}

func (b *Business) SetContact(phone string) {
	b.contact.phone = phone
}

func (b *Business) SetAddress(address string) {
	b.address = address
}

func (b *Business) PrintInfo() {
	fmt.Printf("Name: %s %s %s\n", b.name.fName, b.name.mName, b.name.lName)
	fmt.Printf("Contact: %s\n", b.contact.phone)
	fmt.Printf("Address: %s\n", b.address)
	fmt.Printf("Date of Birth: %s\n", b.dob)
}

// getter function

func (b *Business) GetName() string {
	return fmt.Sprintf("%s %s %s", b.name.fName, b.name.mName, b.name.lName)
}

func (b *Business) GetContct() string {
	return fmt.Sprintf("%s", b.contact.phone)
}
