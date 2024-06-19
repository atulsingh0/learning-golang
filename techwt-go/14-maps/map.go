package main

import "fmt"

func main() {
	//var mp map[string]int
	//mp["one"] = 1
	//mp["two"] = 2
	//mp["three"] = 3

	var mp map[string]int = map[string]int{
		"one":    1,
		"fruits": 4,
		"light":  3,
		"lang":   7,
	}
	fmt.Println("my first map:", mp)

	// another way
	mp2 := make(map[string]string)
	// adding the value
	mp2["one"] = "data"
	mp2["delta"] = "beta"

	fmt.Println("my second map:", mp2)
	fmt.Println("Value at key 'one':", mp2["one"])

	// delete the key
	delete(mp2, "beta") // no error for not existing key
	delete(mp2, "one")

	fmt.Println("my second map after deletion:", mp2)

	// check if key exits

	if val, ok := mp2["gamma"]; ok {
		fmt.Println("Key 'gamma' exists", val)
	} else {
		fmt.Println("Key 'gamma' does not exists")
	}

	fmt.Println("Keys in Map mp", len(mp))

}
