package main

import (
	"fmt"
	"strings"
)

func greet() {
	fmt.Println("Hello, World!")
}

func greetings(name string) {
	fmt.Println("Hello,", name)
}

func mapping(names []string, f func(string)) {
	for _, v := range names {
		f(v)
	}
}

func getSquare(num int) int {
	return num * num
}

func retFunc(num int) func(int) {
	return func(number int) {
		fmt.Println(num * number)
	}
}

func getInitials(name string) (string, string) {
	upperName := strings.Split(strings.ToUpper(name), " ")

	data := []string{}
	for _, name := range upperName {
		data = append(data, name[:1])
	}

	if len(data) > 1 {
		return data[0], data[1]
	}

	return data[0], "_"
}

// --------------

func main() {

	greet()
	greetings("tibbet")

	// calling getLength
	nameList := []string{"enum", "delta", "echo", "bravo", "charlie", "alpha"}

	mapping(nameList, func(name string) { fmt.Println(name, len(name)) })
	mapping(nameList, greetings)

	sqOf3 := getSquare(3)
	fmt.Println("Sq of 3:", sqOf3)

	funcOf4 := retFunc(4)
	funcOf4(2)

	fmt.Println(getInitials("alfred nobel"))
	fmt.Println(getInitials("madam Query"))

}
