package main

import "github.com/gin-gonic/gin"

/*
❯ curl -i "http://localhost:3000/users/adam/28"

HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8
Date: Tue, 08 Aug 2023 22:27:30 GMT
Content-Length: 45

{
    "age": "28",
    "msg": "Hello adam!"
}
*/

func main() {

	rt := gin.Default()

	rt.GET("/users/:name/:age", func(c *gin.Context) {

		c.IndentedJSON(200, gin.H{
			"msg": "Hello " + c.Param("name") + "!",
			"age": c.Param("age"),
		})
	})

	rt.Run(":3000")
}
