package main

import (
	"fmt"
	"os"
)

func main() {

	fh, err := os.Stat("/etc/hosts")
	// Error check
	if err != nil {
		fmt.Println(err.Error())
	}

	// Fetching Stats

	fmt.Println("File Name:", fh.Name())
	fmt.Println("File Size (byte):", fh.Size())
	fmt.Println("File Permissions:", fh.Mode())
	fmt.Printf("File Permissions (octal): %#o\n", fh.Mode().Perm())
	fmt.Println("File, Is it a direcotory?:", fh.IsDir())
	fmt.Println("File, Is it a regular file?:", fh.Mode().IsRegular())
	fmt.Println("File last modification time:", fh.ModTime())
}
