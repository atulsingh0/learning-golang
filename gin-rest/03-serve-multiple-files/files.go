package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, "Hello to Gin world!")
	})

	router.StaticFS("/files", http.FileSystem(http.Dir("./data")))

	router.Run(":3000")
}
