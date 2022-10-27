// Go PlayGround - https://go.dev/play/p/txoyi7cyyW_e
// Formats - https://pkg.go.dev/fmt#hdr-Printing

package main

import "fmt"

func main() {

	a, b := 12, 378
	c, d := 23.871, 5479.00341

	fmt.Println(a, b, c, d)

	fmt.Printf("%d %d\n", a, b) // print as integer
	fmt.Printf("%x %x\n", a, b) // print as hexa decimal value

	fmt.Printf("%b %b\n", a, b)   // print as binary value
	fmt.Printf("%X %X\n", a, b)   // print as hexa decimal value ( letters will be caps)
	fmt.Printf("%#x %#x\n", a, b) // print as hexa decimal with marker
	fmt.Printf("%#b %#b\n", a, b) // print as binary with marker

	fmt.Println()

	fmt.Printf("%f %f\n", c, d)         // default prints the 6 char after decimal
	fmt.Printf("%.2f %.2f\n", c, d)     //2 decimal
	fmt.Printf("%9.2f %9.2f\n", c, d)   //total 9 char = 6 before decimal + 1 decimal + 2 after decimal
	fmt.Printf("%09.2f %09.2f\n", c, d) // filled the 0s
	fmt.Printf("%-9.2f %-9.2f\n", c, d) // Left oriented

}
