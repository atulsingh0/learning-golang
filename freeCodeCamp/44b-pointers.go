package main

import "fmt"

func main() {

	user1 := "Divya"

	admin := &user1

	fmt.Println("Pointer admin and it's value:", admin, *admin)

	user2 := "Ashley"
	admin = &user2

	fmt.Println("Pointer admin and it's value now:", admin, *admin)

	admin2 := admin
	fmt.Println("Pointer admin2:", *admin2)

	user3 := *admin2
	*admin2 = "Priya"

	fmt.Println("admin & admin2 value now:", *admin, *admin2)
	fmt.Println("User3:", user3)
}
