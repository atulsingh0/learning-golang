package main

import (
	"crypto/rand"
	"crypto/rsa"
	"fmt"
	"log"
)

func main() {

	privKey, pubKey := GenerateKeyPair(2048)

	//fmt.Printf("RSA Private Key (%d): %v\n", privKey.N.BitLen(), privKey)
	//fmt.Printf("RSA Public Key: %v", pubKey)

	fmt.Println("\nKeys is human readable format")
	fmt.Printf("Private Key:\n%v", privKey.D)
	fmt.Printf("\nPublic Key:\n%v", pubKey.N)

}

func GenerateKeyPair(bits int) (*rsa.PrivateKey, *rsa.PublicKey) {
	privKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		log.Fatal(err)
	}
	return privKey, &privKey.PublicKey
}
