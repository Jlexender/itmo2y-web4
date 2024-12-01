package handlers

import (
	"github.com/gin-gonic/gin"
	"web4/internal/models"
	"web4/internal/services"
	"web4/internal/util"
)

type RegisterRequest struct {
	Name string `json:"name"`
	Pass string `json:"pass"`
}

type RegisterResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Token   string `json:"token"`
}

func Register(c *gin.Context) {
	var req RegisterRequest
	err := c.BindJSON(&req)

	if err != nil {
		c.JSON(400, RegisterResponse{
			Error:   true,
			Message: "Неправильный формат запроса",
		})
		return
	}

	if !util.IsCredential(req.Name) || !util.IsCredential(req.Pass) {
		c.JSON(400, RegisterResponse{
			Error:   true,
			Message: "Данные введены неверно",
		})
		return
	}

	if services.UserExists(req.Name) {
		c.JSON(400, RegisterResponse{
			Error:   true,
			Message: "Пользователь уже зарегистрирован",
		})
		return
	}

	hash := services.Sha256(req.Pass)
	user := models.User{
		Name: req.Name,
		Hash: hash,
	}

	services.AddUser(user)

	c.JSON(200, RegisterResponse{
		Error:   false,
		Message: "Успешная регистрация",
	})
}
