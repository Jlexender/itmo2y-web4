package routes

import (
	"github.com/gin-gonic/gin"
	"web4/internal/handlers"
)

func Init(r *gin.Engine) {
	base := r.Group("/api/auth")

	base.POST("/register", handlers.Register)
	base.POST("/login", handlers.Login)

}
