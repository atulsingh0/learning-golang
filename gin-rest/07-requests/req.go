package main

import "github.com/gin-gonic/gin"

func main() {

	rt := gin.Default()

	rt.GET("/hi/*rest", func(c *gin.Context) {

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

	rt.Run(":3000")
}
