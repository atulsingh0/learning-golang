package simplemath

import (
	"errors"
	"math"
)

// add function
func Add(p1, p2 float64) float64 {
	return p1 + p2
}

// divide function, multiple return value
func Divide(p1, p2 float64) (float64, error) {
	if p2 == 0 {
		return math.NaN(), errors.New("Can not divide by zero")
	}
	return p1 / p2, nil
}

// variadic functions
func Sum(values ...float64) float64 {
	total := 0.0

	for _, item := range values {
		total += item
	}
	return total
}

// variadic var should be the last input in function
func Mul(p1, p2, p3 float64, p ...float64) float64 {
	out := p1 * p2 * p3

	for _, item := range p {
		out *= item
	}
	return out
}

// named return variable
func Interest(amt float64, rate float64, yr float64) (interest float64) {

	interest = amt * rate * yr / 100 // No need to initialize

	return
}
