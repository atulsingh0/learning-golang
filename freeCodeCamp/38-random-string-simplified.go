package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

const (
	letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-?_"
)

func randomString(num int) string {

	// Creating an array having length <num>
	var randomStr []byte = make([]byte, num)

	for i := 0; i < num; i++ {

		// Generating an random index which varies from 0 to len(letters)
		idx, _ := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		randomStr[i] = letters[idx.Int64()]
	}
	return string(randomStr)
}

func main() {
	fmt.Println(randomString(32))
}
