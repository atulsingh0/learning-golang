package main

import (
	"fmt"
	"html/template"
	"log"
	"os"
)

func main() {

	tmpl, err := template.New("test").Parse("<h1>Hello {{.}}!</h2>")
	if err != nil {
		log.Fatal(err)
	}

	err = tmpl.Execute(os.Stdout, `Gophers`)
	if err != nil {
		return
	}

	fmt.Println()
	// another example
	// Accessing Data Structures

	user := struct {
		Name string
		Age  int
	}{
		Name: "Gopher",
		Age:  25,
	}

	tmpl, err = template.New("us").Parse("<h1>Hello {{.Name}}!</h1></br><b>Age: {{.Age}}</b>")
	if err != nil {
		log.Fatal(err)
	}

	err = tmpl.Execute(os.Stdout, user)
	if err != nil {
		return
	}

}
