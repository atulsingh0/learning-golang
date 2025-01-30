package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := createRouter()

	log.Fatal(r.Run(":3000"))
}

func createRouter() *gin.Engine {
	rt := gin.Default()

	rt.POST("/", func(c *gin.Context) {
		pass := c.DefaultPostForm("password", "missing")
		user := c.PostForm("username")
		c.String(200, fmt.Sprintf("username: %s, password: %s", user, pass))
	})

	return rt
}
