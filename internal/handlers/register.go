package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
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
		c.JSON(http.StatusBadRequest, RegisterResponse{
			Error:   true,
			Message: "Неправильный формат запроса",
		})
		return
	}

	if !util.IsCredential(req.Name) || !util.IsCredential(req.Pass) {
		c.JSON(http.StatusBadRequest, RegisterResponse{
			Error:   true,
			Message: "Данные введены неверно",
		})
		return
	}

	if services.UserExists(req.Name) {
		c.JSON(http.StatusBadRequest, RegisterResponse{
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

	token, err := util.GenerateJWT(user.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, RegisterResponse{
			Error:   true,
			Message: "Ошибка при генерации токена",
		})
		return
	}

	c.JSON(http.StatusOK, RegisterResponse{
		Error:   false,
		Message: "Успешная регистрация",
		Token:   token,
	})
}
