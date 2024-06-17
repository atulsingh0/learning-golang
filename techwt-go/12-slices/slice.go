package main

import "fmt"

func main() {

	no := [8]int{2, 1, 4, 8, 7, 2, 1, 2}

	fmt.Println("no, cap(no), len(no):", no, cap(no), len(no))

	s := no[:]
	fmt.Println("Slice from array 'no':", s)
	fmt.Println("cap(s), len(s):", cap(s), len(s))

	s1 := no[:3]
	fmt.Println("Slice s1:", s1)
	fmt.Println("cap(s1), len(s1):", cap(s1), len(s1))

	// modify the slice
	// change both array and no
	s1[0] = 55
	fmt.Println("Slice s1 and Array no:", s1, no)

	s1 = append(s1, 4)
	fmt.Println("Slice s1 and Array no:", s1, no)

	// beware when making any change to slice made
	// from another array

	// we can reslice a slice
	fmt.Println("re-slice of s1:", s1[1:5])

	// creating slice
	var data []int = []int{2, 3, 4}
	beta := []string{"hi", "how", "are"}

	fmt.Println("data slice:", data)
	fmt.Println("beta slice, cap(beta), len(beta):", beta, cap(beta), len(beta))

	// appending a slice will create a new slice
	fmt.Println("Appending 'you' and '?' in slice beta", append(beta, "you", "?"))
	fmt.Println("Let's reprint beta", beta)

	// to persist we have to store back the new slice
	beta = append(beta, "you", "?")
	fmt.Println("Now reprint beta", beta)
	fmt.Println("cap(beta), len(beta):", cap(beta), len(beta))

	// capacity of slice increase based on requirement
	beta = append(beta, "I", "am", "learning", "go")
	fmt.Println("cap(beta), len(beta):", cap(beta), len(beta))

	// creating a slice from another slice
	gamma := beta[:4]
	fmt.Println(gamma, cap(gamma), len(gamma))

	// Make a copy (deep copy)
	delta := make([]string, len(gamma))
	_ = copy(delta, gamma)

	fmt.Println(delta, cap(delta), len(delta))

	delta[0] = "Oy"
	fmt.Println(delta, cap(delta), len(delta))
	fmt.Println(gamma, cap(gamma), len(gamma))

}
