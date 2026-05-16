package routes

import (
	"github.com/gin-gonic/gin"
)

func Hello(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello, Ved!",
	})
}
func HelloUser(c *gin.Context) {
	name := c.Param("name")
	c.JSON(200, gin.H{
		"message": "Hello, " + name + "!",
	})
}
