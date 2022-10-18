package main

import (
	"encoding/json"
	"fmt"
	"log"
)

const data = `{
	  "id": 1,
	  "name": "Leanne Graham",
	  "username": "Bret",
	  "email": "Sincere@april.biz",
	  "phone": "1-770-736-8031 x56442",
	  "website": "hildegard.org",
	  "company": {
		"name": "Romaguera-Crona",
		"catchPhrase": "Multi-layered client-server neural-net",
		"bs": "harness real-time e-markets"
		}
	}`

type jData struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	EmailId  string `json:"email"`
}

func main() {

	fmt.Printf("%T\n", data)

	// Unarshaling and mapping to struct
	user := jData{}
	err := json.Unmarshal([]byte(data), &user)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(user)
	fmt.Println(user.Name)
	fmt.Println(user.EmailId)

	// Marshaling

	user2 := jData{
		Id:      23,
		Name:    "Apna",
		EmailId: "apna@yandex.com",
		// Username: "apnadesi",
	}
	fmt.Println(user2)

	user2Json, err := json.Marshal(user2)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(user2Json))

}
