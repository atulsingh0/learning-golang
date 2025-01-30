package main

import (
	"log"
	"math/rand"
	"os"
	"strings"
)

func getRandomPic() string {
	// read the file
	data, err := os.ReadFile("pic.list")
	if err != nil {
		return err.Error()
	}
	// parsing the file into slice
	picsArr := strings.Split(string(data), "\n")
	// get random pic
	pic := randomPic(picsArr)
	log.Println("selecting: ", pic)

	// write the pic to the response
	return pic
}

func randomPic(arr []string) string {
	return arr[rand.Intn(len(arr))]
}
