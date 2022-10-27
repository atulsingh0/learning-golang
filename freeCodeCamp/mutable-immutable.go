package main

import "fmt"

func main() {

	i := 4
	j := "Hi"
	k := 2.3

	p := i
	q := j
	r := k

	i = 5
	j = "Bye"

	fmt.Println(i, j, k)
	fmt.Println(p, q, r)

	arr := [4]int{1,2,3,4}
	brr := arr
	arr[2] = 99

	fmt.Println(arr, brr)


		// Struct
		type emp struct {
			Name string
			Job string
		}

		emp1 := emp{
			Name: "Remo",
			Job: "Manager",
		}

		emp9 := emp1

		emp1.Name = "Peter"

		fmt.Println(emp1, emp9)


	// Mutable

	// Slices
	crr := arr[:]
	arr[2] = 100
	fmt.Println(arr, crr)

	// Maps
	kv := map[int]int{1:1, 2:4, 3:9, 4:16}

	keyval := kv
	kv[3] = 27
	fmt.Println(kv, keyval)


}