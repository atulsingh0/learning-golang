package main

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

/*
curl -X POST "http://localhost:3000/users" -d '{"age":28, "firstname": "Akash" }'

{"Age":"28","msg":"Hi Akash"}
*/

type User struct {
	Name string `form:"firstname" json:"firstname"`
	Age  int    `form:"age"  json:"age"`
}

func main() {

	rt := gin.Default()

	rt.POST("/users", func(c *gin.Context) {
		var user User

		err := c.BindJSON(&user)
		fmt.Println(user)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{
			"Msg": "Hi " + user.Name,
			"Age": strconv.Itoa(user.Age),
		})

	})

	apiGrp := rt.Group("/api")
	apiGrp.POST("/users", func(c *gin.Context) {
		var user User

		// To Read form data, we have to use Bind method, rather than BindJSON
		err := c.Bind(&user)
		fmt.Println(user)
		if err != nil {
			c.JSON(400, gin.H{"error from /api/users": err.Error()})
			return
		}

		c.JSON(200, gin.H{
			"msg": "Hello " + user.Name,
			"age": strconv.Itoa(user.Age),
		})

	})

	rt.Run(":3000")
}
