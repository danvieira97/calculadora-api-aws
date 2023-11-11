package main

import "github.com/gin-gonic/gin"

type Numbers struct {
	Number1 float64 `json:"number1"`
	Number2 float64 `json:"number2"`
}

func main() {
	r := gin.Default()

	r.POST("/soma", func(c *gin.Context) {
		var numbers Numbers
		if err := c.ShouldBindJSON(&numbers); err != nil {
			c.JSON(400, gin.H{
				"error": "Invalid payload",
			})
			return
		}
		c.JSON(200, gin.H{
			"result": numbers.Number1 + numbers.Number2,
		})
	})

	r.POST("/subtracao", func(c *gin.Context) {
		var numbers Numbers
		c.BindJSON(&numbers)
		c.JSON(200, gin.H{
			"result": numbers.Number1 - numbers.Number2,
		})
	})

	r.POST("/multiplicacao", func(c *gin.Context) {
		var numbers Numbers
		c.BindJSON(&numbers)
		c.JSON(200, gin.H{
			"result": numbers.Number1 * numbers.Number2,
		})
	})

	r.POST("/divisao", func(c *gin.Context) {
		var numbers Numbers
		c.BindJSON(&numbers)
		c.JSON(200, gin.H{
			"result": numbers.Number1 / numbers.Number2,
		})
	})

	r.Run()
}
