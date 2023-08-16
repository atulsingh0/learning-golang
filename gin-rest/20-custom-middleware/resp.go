package main

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {

	// Responding with File data
	r := gin.Default()

	// Register global middleware
	r.Use(myMiddleWare)

	// Basic Auth middleware
	r.Use(gin.BasicAuth(gin.Accounts{"admin": "password"}))

	r.GET("/reports", myMiddleWare, func(c *gin.Context) { // Registering middleware on route

		// Direct Serve the local file
		c.File("./go.mod")
	})

	r.GET("/reports2", myMiddleWare, func(c *gin.Context) {

		// Serving File from FileSystem
		fs := gin.Dir("./", true)
		c.FileFromFS("./go.sum", fs)
	})

	log.Fatal(r.Run(":3000"))

}

func myMiddleWare(c *gin.Context) {
	t := time.Now()
	c.Next()
	latency := time.Since(t)
	log.Print("processing time taken to process call", c.Request.URL, latency)

}
