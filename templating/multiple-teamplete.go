package main

import (
	"log"
	"os"
	"text/template"
)

type Article struct {
	Name   string
	Author string
}

type Articles []Article

func main() {

	articles := []Article{
		{
			Name:   "How to build a code completion bot",
			Author: "Cody",
		},
		{
			Name:   "Another article",
			Author: "Jane Doe",
		},
	}

	// fmt.Println(article)

	// creating single line template
	// Struct and Map assignment is same
	tmpl, err := template.New("sample").Parse(`
	{{ define "article" }}
		<p>Book "{{.Name}}" is written by "{{.Author}}"</p>
	{{ end }}

	{{ define "articles" }}
		{{ range . }}
			{{ template "article" . }}
		{{ else }}
			<p> No published articles yet </p>
		{{ end }}
	{{ end }}

	{{ template "articles" . }}
`)
	if err != nil {
		log.Fatal(err)
	}
	err = tmpl.Execute(os.Stdout, articles)
	if err != nil {
		log.Fatal(err)
	}
}
