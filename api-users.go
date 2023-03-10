package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {

	fmt.Println("Test API base")

	r := gin.Default()
	r.GET("/users", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Page des utilisateurs",
		})
	})

	r.Run()

}
