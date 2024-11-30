package main

import (
	"github.com/gin-gonic/gin"
	"web4/internal/routes"
)

func main() {
	r := gin.Default()
	routes.Init(r)
	err := r.Run()

	if err != nil {
		panic(err)
	}
}
