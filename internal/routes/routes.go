package routes

import (
	"github.com/gin-gonic/gin"
	"log"
	"web4/internal/handlers"
)

func Init(r *gin.Engine) {
	log.Printf("Starting routes initialization")

	err := r.SetTrustedProxies([]string{"127.0.0.1"})
	if err != nil {
		panic(err)
	}
	base := r.Group("/api/auth")

	base.POST("/register", handlers.Register)
	base.POST("/login", handlers.Login)

	r.GET("/api/user/:name", handlers.GetUser)
}
