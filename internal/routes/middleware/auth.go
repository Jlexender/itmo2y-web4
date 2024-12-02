package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"web4/internal/util"
)

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
			c.Abort()
			return
		}

		username, err := util.ParseJWT(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		log.Printf("User %s has authenticated\n", username)
		c.Set("username", username)
		c.Next()
	}
}
