package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

// {
// 	"SiteName": "My Cat Blog",
// 	"SiteUrl": "www.mycatblog.com",
// 	"Database": {
// 		"Name": "cats",
// 		"Host": "localhost",
// 		"Port": 3306,
// 		"Username": "user1",
// 		"Password": "Password1"
// 	}
// }

type Database struct {
	Name     string
	Host     string
	Port     int16
	Username string
	Password string
}

type Config struct {
	SiteName string
	SiteUrl  string
	Database Database
}

func main() {

	// reading the file
	f, _ := os.Open("config.json")
	defer f.Close()
	data, err := io.ReadAll(f)
	if err != nil {
		panic(err)
	}
	fmt.Print(string(data))

	// Mapping data to Json
	config := Config{}
	err = json.Unmarshal(data, &config)
	if err != nil {
		panic(err)
	}

	// Fetching the details
	fmt.Printf("Site : %s (%s)\n", config.SiteName, config.SiteUrl)

	fmt.Printf("DB : mysql://%s:%s@%s:%d/%s",
		config.Database.Username, config.Database.Password, config.Database.Host, config.Database.Port, config.Database.Name)

}
