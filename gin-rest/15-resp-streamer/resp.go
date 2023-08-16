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
