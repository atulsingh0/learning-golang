package main

import (
	"fmt"
	"math"
)

func main() {

	// typed numeric
	// untyped numetic
	// can not compare mismatched types int and int32

	var (
		alpha int   = 5
		beta  int32 = 21
	)

	result := int32(alpha) == beta

	fmt.Println(result)

	fmt.Println(Uint8FromtInt(200))
	fmt.Println(Uint8FromtInt(200000))

}

func Uint8FromtInt(x int) (uint8, error) {
	if 0 <= x && x <= math.MaxUint8 {
		return uint8(x), nil
	}

	return 0, fmt.Errorf("%d is out of range of uint8\n", x)
}
