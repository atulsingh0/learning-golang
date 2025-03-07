package main

import (
	"errors"
	"fmt"
)

type customError struct {
	msg string
}

func main() {

	err := errors.New("this is a test")

	err2 := conntectAPI()

	// Compare the error
	// an equality check
	if errors.Is(err, err2) {
		fmt.Println("Error is the same")
	} else {
		fmt.Println("Error is not the same")
	}

	// Specific error type check
	// check if return error type is same or not
	var err3 *customError = &customError{
		msg: "connection refused 2nd time",
	}

	err4 := connect()
	if errors.As(err4, &err3) {
		fmt.Println("Error is the same")
	} else {
		fmt.Println("Error is not the same")
	}

}

func conntectAPI() error {
	return errors.New("connection refused")
}

func connect() error {
	return &customError{
		msg: "connection refused",
	}
}
func (e *customError) Error() string {
	return e.msg
}
