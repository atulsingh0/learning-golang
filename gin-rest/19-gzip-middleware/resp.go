package main

import (
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

func main() {

	// Responding with File data
	r := gin.Default()

	r.Use(gzip.Gzip(gzip.DefaultCompression))
	// Basic Auth middleware
	r.Use(gin.BasicAuth(gin.Accounts{"admin": "password"}))

	r.GET("/reports", func(c *gin.Context) {

		// Direct Serve the local file
		c.File("./go.mod")
	})

	r.GET("/reports2", func(c *gin.Context) {

		// Serving File from FileSystem
		fs := gin.Dir("./", true)
		c.FileFromFS("./go.sum", fs)
	})

	r.GET("/msg", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"msg": "This is golang",
		})
	})

	r.GET("/data", func(c *gin.Context) {
		// Opening the file
		fh, err := os.Open("./go.mod")
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
		}
		defer fh.Close()

		// Reading the content
		data, err := io.ReadAll(fh)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
		}

		c.Data(http.StatusOK, "text/plain", data)
	})

	r.GET("/down", func(c *gin.Context) {
		// Opening the file
		fh, err := os.Open("./go.mod")
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
		}
		defer fh.Close()

		// Reading the file stats
		fs, err := fh.Stat()
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
		}

		c.DataFromReader(http.StatusOK,
			fs.Size(),
			"text/plain",
			fh,
			map[string]string{
				"Content-Disposition": `attachment; filename="` + fs.Name() + `"`,
			})
	})

	log.Fatal(r.Run(":3000"))

}
