package main

import (
	"fmt"
	"strings"
	"time"
)

func main(){

	inp := "This is calling of your future."
	
	fmt.Println("Current Time: ", time.Now())
	fmt.Println("Current Day: ", time.Now().Day())
	fmt.Println("Current Month: ", time.Now().Month())
	fmt.Println("Current Year: ", time.Now().Year())


	replacer := strings.NewReplacer(" ", "_")
	fmt.Println(replacer.Replace(inp))

}