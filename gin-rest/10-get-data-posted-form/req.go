package main

import "github.com/gin-gonic/gin"

/*
curl -X POST "http://localhost:3000/users" --form 'name="adam"' --form 'age="28"'

{
    "Age": "28",
    "User": "adam"
}
*/

func main() {

	rt := gin.Default()

	rt.POST("/users", func(c *gin.Context) {

		c.IndentedJSON(200, gin.H{
			"User": c.PostForm("name"),
			"Age":  c.PostForm("age"),
		})
	})

	rt.Run(":3000")
}
