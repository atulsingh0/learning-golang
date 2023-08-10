package main

import (
	"github.com/gin-gonic/gin"
)

func main() {

	// Responding with File data

	r := gin.Default()

	r.GET("/reports", func(c *gin.Context) {

		// Direct Serve the local file
		c.File("./go.mod")
	})

	r.GET("/reports2", func(c *gin.Context) {

		// Serving File from FileSystem
		fs := gin.Dir("./", true)
		c.FileFromFS("./go.sum", fs)
	})

	r.GET("/reports3", func(c *gin.Context) {

		// Serving File from FileSystem

		c.FileAttachment("./resp.go", "goProgram.go")
	})

	r.Run(":3000")
}
