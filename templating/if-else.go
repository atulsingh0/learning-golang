package main

import (
	"log"
	"os"
	"text/template"
)

type Article struct {
	Name   string
	Author string
	Draft  bool
}

type Articles []Article

func main() {

	articles := []Article{
		{
			Name:   "How to build a code completion bot",
			Author: "Cody",
			Draft:  true,
		},
		{
			Name:   "Another article",
			Author: "Jane Doe",
			Draft:  false,
		},
	}

	// fmt.Println(article)

	// creating single line template
	// Struct and Map assignment is same
	tmpl, err := template.New("sample").Parse(`
	{{ range . }}
		<p> Book "{{.Name}}" is written by "{{.Author}}" {{ if .Draft }} (Draft) {{ end }} </p>
	{{ else }}
		<p> No published articles yet </p>
	{{ end }}
`)
	if err != nil {
		log.Fatal(err)
	}
	err = tmpl.Execute(os.Stdout, articles)
	if err != nil {
		log.Fatal(err)
	}
}
