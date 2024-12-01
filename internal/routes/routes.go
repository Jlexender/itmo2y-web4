package routes

import (
	"github.com/gin-contrib/cors"
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

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders: []string{"Origin"},
	}))

	base := r.Group("/api/auth")
	base.POST("/register", handlers.Register)
	base.POST("/login", handlers.Login)

	r.GET("/api/user/:name", handlers.GetUser)
}
