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
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: false,
	}))

	base := r.Group("/api/auth")
	base.POST("/register", handlers.Register)
	base.POST("/login", handlers.Login)

	r.GET("/api/user/:name", handlers.GetUser)
	r.POST("/api/check", handlers.CheckIfInArea)
}
