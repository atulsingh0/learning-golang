package main

import "fmt"

type Printer interface {
	Display() string
}

type item struct {
	Printer
}

func (i *item) Display() string {
	return i.Printer.Display()
}

func main() {

	books := []Book{
		{Title: "The Hitchhiker's Guide to the Galaxy", Author: "Douglas Adams"},
		{Title: "The Lord of the Rings", Author: "J.R.R. Tolkien"},
		{Title: "The Hobbit", Author: "J.R.R. Tolkien"},
		{Title: "The Silmarillion", Author: "J.R.R. Tolkien"},
	}

	games := []Game{
		{Title: "The Legend of Zelda: Breath of the Wild", Developer: "Nintendo"},
		{Title: "Super Mario Odyssey", Developer: "Nintendo"},
		{Title: "The Witcher 3: Wild Hunt", Developer: "CD Projekt Red"},
	}

	//for _, book := range books {
	//	fmt.Println(book.Display())
	//}
	//
	//for _, game := range games {
	//	fmt.Println(game.Display())
	//}

	var it []item
	for _, book := range books {
		it = append(it, item{Printer: &book})
	}
	for _, game := range games {
		it = append(it, item{Printer: &game})
	}

	for _, i := range it {
		fmt.Println(i.Display())
	}

}
