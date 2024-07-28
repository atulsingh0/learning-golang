package main

import (
	"html/template"
	"log"
	"os"
)

func main() {

	tmpl, err := template.New("test").Parse(`{{ . | printf "Price: $%.2f CAD" }}`)

	if err != nil {
		log.Fatal(err)
	}

	err = tmpl.Execute(os.Stdout, 12.345)
	if err != nil {
		return
	}

}
