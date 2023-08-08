package main

import "github.com/gin-gonic/gin"

func main() {

	rt := gin.Default()

	rt.GET("/hi/:name/:age", func(c *gin.Context) {

		c.IndentedJSON(200, gin.H{
			"msg": "Hello " + c.Param("name") + "!",
			"age": c.Param("age"),
		})
	})

	rt.Run(":3000")
}
