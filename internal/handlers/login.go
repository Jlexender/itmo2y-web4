package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
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
		c.JSON(http.StatusBadRequest, LoginResponse{
			Error:   true,
			Message: "Неправильный формат запроса",
		})
		return
	}

	if !util.IsCredential(req.Name) || !util.IsCredential(req.Pass) {
		c.JSON(http.StatusBadRequest, RegisterResponse{
			Error:   true,
			Message: "Данные указаны неверно",
		})
		return
	}

	hash := services.Sha256(req.Pass)
	_, user := services.GetUser(req.Name)

	if user.Hash != hash {
		c.JSON(http.StatusBadRequest, RegisterResponse{
			Error:   true,
			Message: "Данные указаны неверно",
		})
		return
	}

	token, err := util.GenerateJWT(req.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, LoginResponse{
			Error:   true,
			Message: "Ошибка при генерации токена",
		})
		return
	}

	c.JSON(http.StatusOK, LoginResponse{
		Error:   false,
		Message: "Успешный вход",
		Token:   token,
	})
}
