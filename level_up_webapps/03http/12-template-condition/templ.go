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
		"<h2> {{.AuthorName}}!</h2>" +
			"{{ if .Draft }} is writing {{ else }} wrote {{ end }}" +
			"the book {{.Name}}")

	if err != nil {
		log.Fatal(err)
	}

	article := Article{
		Name:       "Golang is awesome",
		AuthorName: "Asta Vista",
		Draft:      false,
	}
	err = tmpl.Execute(os.Stdout, article)
	if err != nil {
		return
	}
}
