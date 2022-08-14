package main

import (
	"fmt"
	"strings"
)

func main() {

	str1 := "this is me"
	str2 := "Kaise na mujge"
	str3 := "This is ME"

	fmt.Println(compareStringIns(str1, str2))
	fmt.Println(compareStringIns(str1, str3))

}

func compareStringIns(a string, b string) bool {
	if len(a) == len(b) {
		if strings.ToLower(a) == strings.ToLower(b) {
			return true
		}
	}
	return false
}
