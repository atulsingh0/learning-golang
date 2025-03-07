package main

import (
	"fmt"
	"unsafe"
)

// the architecture of the machine I am using is 64-bit (8 bytes)
// 8 bytes for int, 8 bytes for float, 1 byte for bool
// 16 bytes for string
type MyStruct struct { // 86 byte
	Name    string  // 16 byte = 128 bits  (8,8)
	Address string  // 16 byte = 128 bits  (8,8)
	Suspend bool    // 1 byte = 8 bits     (8)
	Age     int     // 8 byte = 64 bits    (8)
	Salary  int64   //  8 byte = 64 bits    (8)
	Height  float32 // 4 byte = 32 bits    (8)
}

// rearragning the struct to make it smaller
type MyStruct2 struct { // 56 byte
	Name    string  // 16 byte = 128 bits  (8,8)
	Address string  // 16 byte = 128 bits  (8,8)
	Suspend bool    // 1 byte = 8 bits     (8/2)
	Height  float32 // 4 byte = 32 bits    (8/2)
	Age     int     // 8 byte = 64 bits    (8)
	Salary  int64   //  8 byte = 64 bits    (8)
}

func main() {

	var s MyStruct

	const infoSize = unsafe.Sizeof(s)
	fmt.Println(infoSize)

	fmt.Println(s)

	var data = MyStruct{
		Name:    "Alex",
		Address: "Austin, TX",
		Suspend: false,
		Age:     30,
	}
	fmt.Println(unsafe.Sizeof(data))

	fmt.Println("Length of Name", len(data.Name), len(s.Name))
	fmt.Println("Size of Name:", unsafe.Sizeof(data.Name), unsafe.Sizeof(s.Name))

	data.Name = "Alexander the Great"
	fmt.Println("now, Size of data struct:", unsafe.Sizeof(data), unsafe.Sizeof(s))
	data.Name = "Alexander the Great" + " " + "Alexander the Great" + " " + "Alexander the Great" +
		" " + "Alexander the Great" + " " + "Alexander the Great" + " " + "Alexander the Great" +
		" " + "Alexander the Great"

	fmt.Println("now, Length of Name", len(data.Name), len(s.Name))
	fmt.Println("now, Size of data struct:", unsafe.Sizeof(data), unsafe.Sizeof(s))

	var p MyStruct2
	fmt.Println("Size of p struct:", unsafe.Sizeof(p))

}
