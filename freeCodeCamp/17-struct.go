package main

import "fmt"

type students struct {
	id       int
	name     string
	age      int
	subjects subs
}

type subs struct {
	subs []string
}

func main() {

	// creating a new student
	std := students{
		id:   2,
		name: "ashley",
		age:  21,
		subjects: subs{
			subs: []string{"math", "eng", "science"},
		},
	}

	fmt.Println(std)
	fmt.Println(std.name)
	fmt.Println(std.subjects)
	fmt.Println(std.subjects.subs)

	// anonymous struct (should be short lived)
	aDoc := struct {
		name string
		sr   int
	}{name: "casey", sr: 23}

	fmt.Println("Anonynous Struct:", aDoc)

	bDoc := aDoc  // makes a true copy
	cDoc := &aDoc // referencing to aDoc

	bDoc.name = "alia"
	fmt.Println(aDoc)
	fmt.Println(bDoc)

	cDoc.name = "julia"
	fmt.Println(aDoc)
	fmt.Println(cDoc)

}
