package main

import "fmt"

func main() {

	// Maps are dictionary
	// Maps support only one type of data for keys and values
	students := map[string]int{
		"adam":   23,
		"priya":  21,
		"diya":   21,
		"ashley": 22,
	}

	fmt.Println(students)
	fmt.Println(students["ashley"])

	data := make(map[int]int)
	fmt.Println(data)

	data = map[int]int{
		2: 4,
		3: 9,
		4: 16,
		5: 25,
		6: 36,
	}
	fmt.Println(data)

	// modify the value
	students["ashley"] = 23
	students["ryan"] = 25
	fmt.Println(students)

	// delete
	delete(students, "ryan")
	fmt.Println(students)
	fmt.Println(students["ryan"]) // check if ryan exists

	// the problem with above syntax is, it does not tell us if value is 0 or key does not exists
	pop, ok := students["ryan"]
	fmt.Println(pop, ok) // ok false means, key does not exists
	pop, ok = students["ashley"]
	fmt.Println(pop, ok)

	// lenght of map
	fmt.Println(len(students))

	// Manipulating a copy of map will modify the original map as well.
	sd := students // reference to student data
	sd["disha"] = 21
	fmt.Println(students)

}
