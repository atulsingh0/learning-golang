package main

import "fmt"

func main() {

	var alpha map[string]string = map[string]string{
		"name": "atul",
		"job":  "SE",
		"city": "toronto",
	}

	fmt.Printf("%v\n", alpha)

	//  CRUD operation

	var beta = make(map[string]string) // create
	fmt.Printf("beta: %v\n", beta)

	beta["name"] = "Priya" //  add
	beta["job"] = "chef"
	beta["city"] = "motegomery"
	beta["age"] = "25"
	beta["org"] = "fine dining"
	fmt.Printf("Add operation: beta: %v\n", beta)

	beta["job"] = "cook" // update
	fmt.Printf("Update operation: beta: %v\n", beta)

	delete(beta, "city") // delete
	fmt.Printf("Delete operation: beta: %v\n", beta)

	delete(beta, "gender") // delete (key does not exist)
	fmt.Printf("Delete operation:  beta: %v\n", beta)

	name := beta["name"]
	fmt.Printf("beta: %v\n", beta)
	fmt.Printf("val: %v\n", name)

	age, status := beta["age"]
	fmt.Printf("beta: %v\n", beta)
	fmt.Printf("age: %v \t status: %v\n", age, status)

	gender, status := beta["gender"]
	fmt.Printf("beta: %v\n", beta)
	fmt.Printf("gender: %v \t status: %v\n", gender, status)

	fmt.Printf("Length of map: %v\n", beta)

	for k, v := range beta {
		fmt.Printf("%v: %v\n", k, v)
	}

	keys := make([]string, 0, len(beta))
	for k := range beta {
		keys = append(keys, k)
	}
	fmt.Printf("Keys in beta: %v\n", keys)

}
