package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {

	var str string = "The time " + time.Now().Format(time.Kitchen) + " is now."
	fmt.Println(str)

	fmt.Printf("The time %s is now.\n", time.Now().Format(time.Kitchen))

	var strng string = fmt.Sprintf("The time %s is now.\n", time.Now().Format(time.Kitchen))
	fmt.Println(strng)

	str_arr := []string{"The time", time.Now().Format(time.Kitchen), "is now."}
	fmt.Println(str_arr)
	fmt.Printf("%s\n\n", strings.Join(str_arr, " "))

	// use strings.Builder to build a string
	// it allows you to build a string without needing to know the size of the string
	// beforehand. It's a bit more expensive to use, but it's useful when you don't know
	// the size of the string you're building.

	var line strings.Builder
	fmt.Println("----------------")
	line.WriteString("The time ")
	line.WriteString(time.Now().Format(time.Kitchen))
	fmt.Println(line.String())
	line.WriteString(" is now.")
	fmt.Println("----------------")
	fmt.Println(line.String())
	line.WriteString(" Ohh, it is continuing.\n")
	fmt.Println(line.String())

	// reset the string builder
	line.Reset()
	fmt.Println(line.String())
	line.WriteString("A new line")

	// strings.Builder implements io.Writer interface
	fmt.Fprint(&line, " and another one.")
	fmt.Fprint(&line, "and a no ")
	fmt.Fprint(&line, 9232)
	fmt.Fprint(&line, " as well.\n")
	fmt.Println(line.String())

}
