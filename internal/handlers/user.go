package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Register(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "world",
	})
}
