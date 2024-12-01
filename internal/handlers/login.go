package handlers

import (
	"github.com/gin-gonic/gin"
	"web4/internal/services"
	"web4/internal/util"
)

type LoginRequest struct {
	Name string `json:"name"`
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
			Message: "Неправильный формат запроса",
		})
		return
	}

	if !util.IsCredential(req.Name) || !util.IsCredential(req.Pass) {
		c.JSON(400, RegisterResponse{
			Error:   true,
			Message: "Данные указаны неверно",
		})
		return
	}

	hash := services.Sha256(req.Pass)
	_, user := services.GetUser(req.Name)

	if user.Hash != hash {
		c.JSON(400, RegisterResponse{
			Error:   true,
			Message: "Данные указаны неверно",
		})
		return
	}

	c.JSON(200, RegisterResponse{
		Error:   false,
		Message: "Успешный вход",
	})
}
