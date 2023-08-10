package main

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

/*
curl -X POST "http://localhost:3000/users" -d '{"age":28, "firstname": "Akash", "role": "admin" }'
{"Age":"28","Msg":"Hi Akash","Role":"admin"}

curl -X POST "http://localhost:3000/users" -d '{ "firstname": "Akash", "role": "admin" }'
{"Age":"0","Msg":"Hi Akash","Role":"admin"}

curl -X POST "http://localhost:3000/users" -d '{ "firstname": "Akash", "role": "emp" }'
{"error":"Key: 'User.Role' Error:Field validation for 'Role' failed on the 'oneof' tag"}
*/

type User struct {
	Name string `json:"firstname" binding:"required"`
	Age  int    `json:"age" binding:"-"`
	Role string `json:"role" binding:"oneof=admin employee"`
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
			"Msg":  "Hi " + user.Name,
			"Age":  strconv.Itoa(user.Age),
			"Role": user.Role,
		})

	})

	rt.Run(":3000")
}
