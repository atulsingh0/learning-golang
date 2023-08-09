package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

/*
curl -X POST "http://localhost:3000/users" -F "year"=2010 -F "amount"=50

{
    "amount": "50",
    "user": "anonymous",
    "year": "2010"
}
*/

func main() {

	rt := gin.Default()

	rt.POST("/users", func(c *gin.Context) {

		c.IndentedJSON(200, gin.H{
			"year":   c.PostForm("year"),
			"amount": c.PostForm("amount"),
			"user":   c.DefaultPostForm("user", "anonymous"),
		})
	})

	log.Fatal(rt.Run(":3000"))
}
