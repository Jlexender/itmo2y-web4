package main

import (
	"github.com/gin-gonic/gin"
	"web4/internal/db"
	"web4/internal/routes"
)

func main() {
	err := db.InitDB()
	if err != nil {
		panic(err)
	}

	r := gin.Default()

	routes.Init(r)
	err = r.Run(":8080")
	if err != nil {
		panic(err)
	}
}
