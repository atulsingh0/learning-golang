package main

import (
	"fmt"
)

func main() {

	// Ordered collection of same type of data
	// Fixed length
	// based on data type, value will be defaulted

	var arr [4]int
	fmt.Println(arr)

	arr[3] = 8
	arr[1] = 2
	fmt.Println(arr)

	var sent [8]string
	sent[0] = "Hi"
	sent[1] = "How"

	fmt.Println(sent)

	no := [5]float64{2.3, 1, 4, 0, 3.1}
	fmt.Println(no)

	// length and capacity
	fmt.Println("len and cap of arr", len(arr), cap(arr))
	var sum1 int
	for i := 0; i < len(arr); i++ {
		sum1 = sum1 + arr[i]
	}
	fmt.Println("total:", sum1)
}

curl --request GET \
--url "https://api.snyk.com/rest/orgs/844e0371-ef50-48c1-a0d1-1dbd652b2982/projects?version=2024-06-10" \
--header "Content-Type: application/vnd.api+json" \
--header "Authorization: token ${SYNK_API_TOKEN}"
