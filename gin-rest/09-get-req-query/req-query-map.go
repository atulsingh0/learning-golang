package main

import "github.com/gin-gonic/gin"

/*
curl -i "http://localhost:3000/hi/test?id=3&user=adam"

HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8
Date: Tue, 08 Aug 2023 22:25:43 GMT
Content-Length: 118

{
    "url-query": {
        "id": [
            "3"
        ],
        "user": [
            "adam"
        ]
    }
}
*/

func main() {

	rt := gin.Default()

	rt.GET("/hi/*rest", func(c *gin.Context) {

		c.IndentedJSON(200, gin.H{
			"url-query": c.Request.URL.Query(),
		})
	})

	rt.Run(":3000")
}
