package main

import "fmt"

func main() {

	fmt.Println("welcome to input formatting")

	/*
		%v - default format
		%s - string
		%d - decimal
		%f - floats
		%g - floating-point numbers
		%b - binary numbers
		%o - octal numbers
		%x - Hexa decimal
		%t - boolean value
		%c - ascii char
	*/

	var id int = 23
	var name string = "disha"
	var intrest float32 = 2.124879
	var enabled bool = true

	fmt.Printf("%v\t%T\n", id, id)
	fmt.Printf("%s\t%T\n", name, name)

	fmt.Printf("%g\t%T\n", intrest, intrest)
	fmt.Printf("%f\t%T\n", intrest, intrest)
	fmt.Printf("%.2f\t%T\n", intrest, intrest)

	fmt.Printf("%t\t%T\n\n", enabled, enabled)

	fmt.Printf("|%f|%f|%f|%f|\n", 23.12, 87.123, 89.1112, 65.1)
	fmt.Printf("|%f|%f|%f|%f|\n\n", 23.1212, 7.123, 89.2, 65.0)

	fmt.Printf("|%4.2f|%4.2f|%4.2f|%4.2f|\n", 23.12, 87.123, 89.1112, 65.1)
	fmt.Printf("|%4.2f|%4.2f|%4.2f|%4.2f|\n\n", 23.1212, 7.123, 89.2, 65.0)

	fmt.Printf("|%8.2f|%8.2f|%8.2f|%8.2f|\n", 23.12, 87.123, 89.1112, 65.1)
	fmt.Printf("|%8.2f|%8.2f|%8.2f|%8.2f|\n\n", 23.1212, 7.123, 89.2, 65.0)

	fmt.Printf("|%-8.2f|%-8.2f|%-8.2f|%-8.2f|\n", 23.12, 87.123, 89.1112, 65.1)
	fmt.Printf("|%-8.2f|%-8.2f|%-8.2f|%-8.2f|\n\n", 23.1212, 7.123, 89.2, 65.0)

	fmt.Printf("|%-8.2f|%-8.2s|%8.2q|%-8.2f|\n", 23.12, "Hi", "Name", 65.1)
	fmt.Printf("|%-8.2f|%8s|%-8.2f|%-8.2f|\n\n", 23.1212, "Maan", 89.2, 65.0)

	output := fmt.Sprintf("|%-8.2f|%8s|%-8.2f|%-8.2f|\n\n", 23.1212, "Maan", 89.2, 65.0)
	fmt.Printf(output)
}
