package main

import (
	"encoding/base64"
	"fmt"
)

func main() {

	string_to_encode := "this is going to be base64 encoded."

	encoded_data := encode(string_to_encode)

	fmt.Println("[Base64] Encoded string: ", encoded_data)
	fmt.Println("[Base64] Decoded string: ", decode(encoded_data))

}

// base64 Encode
func encode(data string) string {
	return base64.StdEncoding.EncodeToString([]byte(data))
}

// base64 Decode
func decode(data string) string {

	out, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		panic(err)
	}
	return string(out)
}
