package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {

	// Responding with File data
	r := gin.Default()

	r.GET("/stream", func(c *gin.Context) {
		// Opening the file
		fh, err := os.Open("/var/log/system.log")
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
		}
		defer fh.Close()
		c.Stream(streamer(fh))
	})

	r.GET("/about", func(c *gin.Context) {

		cookies := c.Request.Cookies()
		c.IndentedJSON(200, gin.H{
			"header":   c.Request.Header,
			"url":      c.Request.URL,
			"query":    c.Request.URL.Query(),
			"path":     c.Request.URL.Path,
			"params":   c.Params,
			"body":     c.Request.Body,
			"method":   c.Request.Method,
			"clientIP": c.ClientIP(),
			"cookies":  cookies,
		})
	})

	r.GET("/users/:name/:age", func(c *gin.Context) {

		c.IndentedJSON(200, gin.H{
			"msg": "Hello " + c.Param("name") + "!",
			"age": c.Param("age"),
		})
	})

	log.Fatal(r.Run(":3000"))

}

func streamer(r io.Reader) func(io.Writer) bool {
	return func(step io.Writer) bool {
		for {
			buf := make([]byte, 1*2^10)
			if _, err := r.Read(buf); err != nil {
				return false
			} else {
				_, err := step.Write(buf)

				// added some delay to observe the behaviour
				time.Sleep(time.Microsecond * 10)
				return err == nil
			}

		}
	}
}
