package main

import (
	"os"
	"text/template"
)

func main() {

	tmpl, _ := template.New("foo").Parse(
		"Price: ${{printf \"%.2f\" .}}\n",
	)
	tmpl.Execute(os.Stdout, 11.234)

	// it can be written as pipeline
	tmpl, _ = template.New("foo").Parse(
		"Price: ${{. | printf \"%.2f\" }}\n",
	)
	tmpl.Execute(os.Stdout, 11.234)

}
