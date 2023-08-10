package main

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	"github.com/go-playground/validator/v10"
)

/*
curl -X POST "http://localhost:3000/number" -d '{"number":7 }'
{"Age":"28","Msg":"Hi Akash","Role":"admin"}
*/

type No struct {
	Number int `json:"number" binding:"required,prime"`
}

var prime validator.Func = func(f1 validator.FieldLevel) bool {
	n := f1.Field().Int()
	if n < 2 {
		return false
	}
	for i := 2; i <= int(n/2); i++ {
		if int(n)%i == 0 {
			return false
		}
	}
	return true
}

func main() {

	rt := gin.Default()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("prime", prime)
	}

	rt.POST("/number", func(c *gin.Context) {
		var num No

		err := c.BindJSON(&num)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{
			"Number": "It is a prime number: " + strconv.Itoa(num.Number),
		})

	})

	rt.Run(":3000")
}
