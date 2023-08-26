package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	parseJSON(`{
		"foo" : 123.01
	}`)

	parseJSON(`{
		"foo" : 12
	}`)

	parseJSON(`{
		"foo" : "bar"
	}`)
	parseJSON(`{
		"foo" : []
	}`)
}

func parseJSON(inp string) {

	data := map[string]interface{}{}
	err := json.Unmarshal([]byte(inp), &data)
	if err != nil {
		panic(err)
	}

	// processing key "foo"

	val, _ := data["foo"]

	switch val.(type) {
	case float64:
		fmt.Printf("Float64: %f\n", val)
	case string:
		fmt.Printf("String: %s\n", val)
	default:
		fmt.Println("foo is something which I dont know")
	}

}
