package main

import (
	"fmt"
	"html/template"
	"os"
)

type user struct {
	Name string
	Age  int
}

// String method is used to convert the user struct to string
// niladic method
func (u user) String() string {
	return fmt.Sprintf("we have %s user with age %d.", u.Name, u.Age)
}

func main() {
	tmpl, err := template.New("test").Parse("<h1>Hello, {{.String}}</h2>")
	u := user{
		Name: "Gopher",
		Age:  25,
	}

	err = tmpl.Execute(os.Stdout, u)
	if err != nil {
		return
	}

}
