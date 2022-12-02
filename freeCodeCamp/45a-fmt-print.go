// Go PlayGround - https://go.dev/play/p/txoyi7cyyW_e
// Formats - https://pkg.go.dev/fmt#hdr-Printing

package main

import "fmt"

func main() {

	a, b := 12, 378
	c, d := 23.871, 5479.00341

	// 12 378 23.871 5479.00341
	fmt.Println(a, b, c, d)

	// 12 378
	fmt.Printf("%d %d\n", a, b) // print as integer

	// c 17a
	fmt.Printf("%x %x\n", a, b) // print as hexa decimal value

	// 1100 101111010
	fmt.Printf("%b %b\n", a, b) // print as binary value

	// C 17A
	fmt.Printf("%X %X\n", a, b) // print as hexa decimal value ( letters will be caps)

	// 0xc 0x17a
	fmt.Printf("%#x %#x\n", a, b) // print as hexa decimal with marker

	// 0b1100 0b101111010
	fmt.Printf("%#b %#b\n", a, b) // print as binary with marker

	fmt.Println()

	// 23.871000 5479.003410
	fmt.Printf("%f %f\n", c, d) // default prints the 6 char after decimal

	// 23.87 5479.00
	fmt.Printf("%.2f %.2f\n", c, d) // 2 decimal

	//     23.87   5479.00
	fmt.Printf("|%9.2f|%9.2f|\n", c, d) //total 9 char = 6 before decimal + 1 decimal + 2 after decimal

	// 000023.87 005479.00
	fmt.Printf("|%09.2f|%09.2f|\n", c, d) // filled the 0s

	// 23.87     5479.00
	fmt.Printf("|%-9.2f|%-9.2f|\n", c, d) // Left oriented

}
