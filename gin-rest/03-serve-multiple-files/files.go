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

	// serve the filesystem
	router.StaticFS("/files", http.FileSystem(http.Dir("./data")))

	// server the local filesystem
	router.Static("/index", "./data")

	router.Run(":3000")
}
