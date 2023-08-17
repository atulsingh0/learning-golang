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

func main() {

	article := Article{
		Name:   "How to build a code completion bot",
		Author: "Cody"}

	// fmt.Println(article)

	// creating single line template
	// Struct and Map assignment is same
	tmpl, err := template.New("sample").Parse("Book \"{{.Name}}\" is written by \"{{.Author}}\"")
	if err != nil {
		log.Fatal(err)
	}
	err = tmpl.Execute(os.Stdout, article)
	if err != nil {
		log.Fatal(err)
	}
}
