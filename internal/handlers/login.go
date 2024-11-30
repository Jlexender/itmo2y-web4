package handlers

import (
	"github.com/gin-gonic/gin"
	"web4/internal/services"
)

type LoginRequest struct {
	User string `json:"user"`
	Pass string `json:"pass"`
}

type LoginResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Token   string `json:"token"`
}

func Login(c *gin.Context) {
	var req LoginRequest
	err := c.BindJSON(&req)

	if err != nil {
		c.JSON(400, LoginResponse{
			Error:   true,
			Message: "Invalid request",
		})
		return
	}

	if req.User == "" || req.Pass == "" {
		c.JSON(400, RegisterResponse{
			Error:   true,
			Message: "Invalid request",
		})
		return
	}

	hash := services.Sha256(req.Pass)
	user := services.GetUser(req.User)

	if user.Hash != hash {
		c.JSON(400, RegisterResponse{
			Error:   true,
			Message: "Invalid credentials",
		})
		return
	}

	c.JSON(200, RegisterResponse{
		Error:   false,
		Message: "User logged in",
	})
}
