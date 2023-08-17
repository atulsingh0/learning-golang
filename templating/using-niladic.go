package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

type Article struct {
	Name   string
	Author string
}

func (a Article) Byline() string {
	return fmt.Sprintf("Name: %s - Author: %s", a.Name, a.Author)
}

func main() {

	article := Article{
		Name:   "How to build a code completion bot",
		Author: "Cody"}

	// fmt.Println(article)

	// creating single line template
	// Struct and Map assignment is same
	tmpl, err := template.New("sample").Parse("{{ .Byline }}")
	if err != nil {
		log.Fatal(err)
	}
	err = tmpl.Execute(os.Stdout, article)
	if err != nil {
		log.Fatal(err)
	}
}
