package main

import (
	"html/template"
	"log"
	"os"
)

type Article struct {
	Name       string
	AuthorName string
	Draft      bool
}

func main() {

	tmpl, err := template.New("test").Parse(
		"{{ range.}}" +
			"<h2> {{.AuthorName}}!</h2>" +
			"{{ if .Draft }} is writing {{ else }} wrote {{ end }}" +
			"the book {{.Name}}" +
			"{{ end }}")

	if err != nil {
		log.Fatal(err)
	}

	articles := []Article{
		Article{
			Name:       "Golang is awesome",
			AuthorName: "Asta Vista",
			Draft:      false,
		},
		Article{
			Name:       "Golang Programming",
			AuthorName: "Mika de",
			Draft:      true,
		}}

	err = tmpl.Execute(os.Stdout, articles)
	if err != nil {
		return
	}
}
