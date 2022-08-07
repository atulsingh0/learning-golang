package main

import (
	"fmt"
	"strings"
)

func main() {

	str1 := "this is me calling func"
	str2 := "this is again me learning go."
	str3 := "this is me calling func"

	if str1 == str2 {
		fmt.Println("str1 and str2 is same")
	} else {
		fmt.Println("str1 and str2 is different")
	}

	if str1 == str3 {
		fmt.Println("str1 and str3 is same")
	} else {
		fmt.Println("str1 and str3 is different")
	}

	if strings.Compare(str1, str3) == 0 {
		fmt.Println("str1 and str3 is same")
	} else {
		fmt.Println("str1 and str3 is different")
	}

	cities := []string{"Toronot", "Alabama", "Delhi", "Agra", "Quebec"}

	for _, city := range cities {
		fmt.Println(strings.Compare("Delhi", city))
	}
}
