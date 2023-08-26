package main

import (
	"os"
	"text/template"
)

type Product struct {
	Price    float64
	Quantity float64
}

func main() {

	tmpl := template.New("foo")

	// registering the template
	tmpl.Funcs(template.FuncMap{"multiply": Multiply})

	// parsing the template
	tmpl, err := tmpl.Parse(
		"Price: ${{ multiply .Price .Quantity | printf \"%.2f\" }}\n",
	)
	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(os.Stdout, Product{
		Price:    25.124,
		Quantity: 89.0123,
	})
	if err != nil {
		panic(err)
	}
}

// let's define a func which process the price before we print it
func Multiply(price, quantity float64) float64 {
	return price * quantity
}
