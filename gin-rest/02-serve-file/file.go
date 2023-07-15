package main

import "github.com/gin-gonic/gin"

func main() {

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, "Hello to Gin world!")
	})

	router.StaticFile("/file", "./test.txt")

	router.Run(":3000")
}
