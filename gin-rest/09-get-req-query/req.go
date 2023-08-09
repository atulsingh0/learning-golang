package main

import (
	"log"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

/*
curl -X GET "http://localhost:3000/query/?user=adam&year=2010&month=2&month=5"

{
    "Months": [
        "2",
        "5"
    ],
    "User": "adam",
    "Year": "2010"
}
*/

func main() {

	rt := gin.Default()

	rt.GET("/query/*rest", func(c *gin.Context) {

		c.IndentedJSON(200, gin.H{
			"User":      c.Query("user"),
			"Year":      c.DefaultQuery("year", strconv.Itoa(time.Now().Year())),
			"Months":    c.QueryArray("month"),
			"url-query": c.Request.URL.Query(),
		})
	})

	log.Fatal(rt.Run(":3000"))
}
