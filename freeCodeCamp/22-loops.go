package main

import "fmt"

func main() {

	// basic for loop
	for i := 0; i < 3; i++ {
		fmt.Println("I am:", i)
	}

	// including anther var
	for k, j := 0, 0; k < 3; k, j = k+1, j+2 {
		fmt.Println("I am:", k, j)
	}

	p := 0
	for p < 4 {
		fmt.Println("Hey, I am:", p)
		p++
	}

	// indefinite loop
	for {
		fmt.Println("I am running from indefinite loop with p:", p)
		if p == 10 {
			fmt.Println("I am broke cause p is 10")
			break
		}
		p++
	}

	// skip the steps with 'continue'
	for ; p <= 20; p++ {
		fmt.Println("Exploring 'continue' with p:", p)
		if p == 17 {
			fmt.Println("I am 17 so skip")
			continue
		}
		fmt.Println("I am p:", p)
	}

	// Nested loops
	for x := 0; x < 5; x++ {
		fmt.Println("With x:", x)
		for y := 0; y <= 5; y++ {
			fmt.Println("x * y:", x*y)
		}
	}

	// Breaking from Nested Loop
outer:
	for x := 0; x <= 4; x++ {
		fmt.Println("With x:", x)
	inner:
		for y := 0; y <= 4; y++ {
			fmt.Println("x * y:", x*y)

			if x*y == 6 {
				fmt.Println("Break from Inner")
				break inner
			}

			if x*y > 15 {
				fmt.Println("Break from Outer")
				break outer
			}
		}
	}

	// for loop for slices/maps
	s := []int{1, 2, 4}

	for _, v := range s {
		fmt.Println(v)
	}

	m := map[string]string{
		"name": "atul",
		"job":  "SE",
	}

	for key, val := range m {
		fmt.Println(key, val)
	}

}
