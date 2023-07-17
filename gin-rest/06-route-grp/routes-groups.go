package main

import "github.com/gin-gonic/gin"

func main() {

	r := gin.Default()

	userGroup := r.Group("/user")

	userGroup.GET("/name", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"name": "Ashley",
		})
	})

	userGroup.GET("/age", func(c *gin.Context) {
		c.String(200, "20")
	})

	r.Run(":3000")
}
