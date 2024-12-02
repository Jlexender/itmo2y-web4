package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"web4/internal/services"
)

func GetUser(c *gin.Context) {
	name := c.Param("name")
	found, user := services.GetUser(name)

	if !found {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   true,
			"message": "User not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"error":   false,
		"message": user,
	})
}
