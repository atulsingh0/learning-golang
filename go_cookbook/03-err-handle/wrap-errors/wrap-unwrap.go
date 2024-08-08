package main

import (
	"errors"
	"fmt"
)

// error Struct which wrap the error Err
type dbConError struct {
	Host string
	Port int
	Err  error
}

func main() {

	// scenario 1
	err1 := errors.New("Error 1")
	err2 := fmt.Errorf("Error 2 is one: %w", err1)

	fmt.Printf("Print err2: %q\n", err2)

	err3 := errors.Unwrap(err2)
	fmt.Printf("Print err3: %q\n", err3)

	// scenario 2
	// utilizing error interface to create custom error
	derr := dbConError{
		Host: "localhost",
		Port: 3306,
		Err:  errors.New("connection refused"),
	}
	fmt.Println("DB Con Error: ", derr.Error())

	err4 := errors.Unwrap(derr)
	fmt.Printf("Print err4: %q\n", err4)

}

// Error implements the error interface
func (e dbConError) Error() string {
	return fmt.Sprintf("Error connecting to %s:%d : %s", e.Host, e.Port, e.Err)
}

// Unwrap implements the interface
func (e dbConError) Unwrap() error {
	return e.Err
}
