package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Name  string
	Age   int
	Class int
	Marks []int
}

func main() {

	// string json
	st := `{
		"Name": "Gopher",
		"Age": 25,
		"Class": 12,
		"Marks": [100, 200, 300]
	}`

	var s Student

	// convert json to Student struct
	// deserialize
	err := json.Unmarshal([]byte(st), &s)
	if err != nil {
		return
	}

	fmt.Println(s.Class)

}
