package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)

	router.Run("localhost:9090")
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "Hi, bye")
}
