package main

import (
	"html/template"
	"log"
	"os"
)

type Book struct {
	Title    string
	Price    float32
	Quantity int
}

func totalPrice(Price float32, Quantity int) float32 {
	return Price * float32(Quantity)
}

func main() {

	tmpl := template.New("test")
	tmpl.Funcs(template.FuncMap{
		"totalPrice": totalPrice,
	})

	tmpl, err := tmpl.Parse(
		"Book : {{ .Title }}\n" +
			"Price: ${{.Price }} CAD\n" +
			"Quantity: {{.Quantity }}\n" +
			"{{ totalPrice .Price .Quantity | printf \"Total Price: $%.2f CAD\" }}\n")

	if err != nil {
		log.Fatal(err)
	}

	err = tmpl.Execute(os.Stdout, Book{
		Title:    "Golang Programming",
		Price:    12.345,
		Quantity: 10,
	})
	if err != nil {
		return
	}

}
