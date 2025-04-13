package main

import "fmt"

type Game struct {
	Title     string
	Developer string
}

func (b *Game) Display() string {
	return fmt.Sprintf("Title: %s\tDeveloper: %s", b.Title, b.Developer)
}
