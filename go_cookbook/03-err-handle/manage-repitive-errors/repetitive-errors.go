package main

import (
	"errors"
	"fmt"
)

func main() {

	a := 4
	out, err := calSquare(a)
	if err != nil {
		fmt.Printf(parseError("calSquare", err))
	} else {
		fmt.Printf("Square of %d: %d\n", a, out)
	}

	a = 101
	out, err = calSquare(a)
	if err != nil {
		fmt.Printf(parseError("calSquare", err))
	} else {
		fmt.Printf("Square of %d: %d\n", a, out)
	}
}

// error parser function
// we can add more error types and parse them here
func parseError(msg string, err error) string {
	return fmt.Sprintf("!! Error: %s : %s", msg, err.Error())
}

func calSquare(x int) (int, error) {
	if x < 0 {
		return 0, errors.New("value must be positive")
	}
	if x > 100 {
		return 0, errors.New("value must be less than 100")
	}
	return x * x, nil
}
