package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	// creating a tmp file
	fh, err := os.OpenFile(os.TempDir()+"test.txt", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	defer fh.Close()

	save(fh, "hello My Name is Lakhan")

	// reading data from file
	data, err := os.ReadFile(fh.Name())
	if err != nil {
		panic(err)
	}
	fmt.Println(fh.Name(), "\n", string(data))

}

// io.Writer is interface
func save(ff io.Writer, data string) error {
	_, err := ff.Write([]byte(data))
	return err
}
