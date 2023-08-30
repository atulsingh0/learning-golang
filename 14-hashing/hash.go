package main

import (
	"crypto/md5"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
)

func main() {

	data := "abra ka dabra"
	fmt.Printf("%s [md5 checksum] : %x\n", data, md5.Sum([]byte(data)))
	fmt.Printf("%s [sha224 (sha256) checksum] : %x\n", data, sha256.Sum224([]byte(data)))
	fmt.Printf("%s [sha256 checksum] : %x\n", data, sha256.Sum256([]byte(data)))
	fmt.Printf("%s [sha224 (sha512) checksum] : %x\n", data, sha512.Sum512_224([]byte(data)))
	fmt.Printf("%s [sha256 (sha512) checksum] : %x\n", data, sha512.Sum512_256([]byte(data)))
	fmt.Printf("%s [sha384 (sha512) checksum] : %x\n", data, sha512.Sum384([]byte(data)))
	fmt.Printf("%s [sha512 checksum] : %x\n", data, sha512.Sum512([]byte(data)))

}
