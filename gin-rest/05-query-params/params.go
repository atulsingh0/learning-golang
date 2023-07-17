package main

import "github.com/gin-gonic/gin"

func main() {

	router := gin.Default()

	router.GET("/hi/:user/*rest", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"msg":  "Hello " + c.Param("user"),
			"rest": c.Param("rest"),
		})
	})

	router.GET("/article/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.String(200, id)
	})

	router.Run(":3000")
}
