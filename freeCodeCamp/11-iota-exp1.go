package main

import "fmt"

const (
	_  = iota             // ignore value 0
	KB = 1 << (10 * iota) // bitwise shift of bytes
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

func main() {

	fileSize := 4000000000.0

	fmt.Printf("KB: %d Byte\n", KB)
	fmt.Printf("%0.2fKB\n\n", fileSize/KB)

	fmt.Printf("MB: %d Byte\n", MB)
	fmt.Printf("%0.2fMB\n\n", fileSize/MB)

	fmt.Printf("GB: %d Byte\n", GB)
	fmt.Printf("%0.2fGB\n", fileSize/GB)

}
