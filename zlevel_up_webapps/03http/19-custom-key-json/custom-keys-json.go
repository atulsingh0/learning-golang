package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Class int    `json:"group"`
	Marks []int  `json:"marks"`
}

func main() {

	// string json
	st := `{
		"Name": "Gopher",
		"Age": 25,
		"Group": 12,
		"Marks": [100, 200, 300]
	}`

	var s Student

	// convert json to Student struct
	// deserialize
	err := json.Unmarshal([]byte(st), &s)
	if err != nil {
		return
	}

	fmt.Println(s.Name)
	fmt.Println(s.Age)
	fmt.Println(s.Class)

}
