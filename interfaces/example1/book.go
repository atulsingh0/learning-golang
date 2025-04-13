package main

import "fmt"

type Book struct {
	Title  string
	Author string
}

func (b *Book) Display() string {
	return fmt.Sprintf("Title: %s\tAuthor: %s", b.Title, b.Author)
}
