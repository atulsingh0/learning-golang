package main

import "github.com/gin-gonic/gin"

func main() {

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, "Hello to Gin world!")
	})

	router.GET("/greek", func(c *gin.Context) {
		c.File("./test.txt")
	})

	router.POST("/hi", func(c *gin.Context) {
		c.String(200, "Hello")
	})

	router.Run(":3000")
}
