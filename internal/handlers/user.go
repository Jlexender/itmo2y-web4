package handlers

import (
	"github.com/gin-gonic/gin"
	"web4/internal/services"
)

func GetUser(c *gin.Context) {
	name := c.Param("name")
	found, user := services.GetUser(name)

	if !found {
		c.JSON(404, gin.H{
			"error":   true,
			"message": "User not found",
		})
		return
	}

	c.JSON(200, gin.H{
		"error":   false,
		"message": user,
	})
}
