package main

import (
	"bytes"
	"fmt"
)

func main() {

	var buf bytes.Buffer // bytes.Buffer implements io.Writer

	fmt.Fprintf(&buf, "Hello, %s", "world")
	s := buf.String()
	fmt.Println(s)

}
