package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"web4/internal/config"
	"web4/internal/db"
	"web4/internal/routes"
)

func main() {
	err := config.InitConfig("config/config.json")
	cfg := config.GetConfig()
	if err != nil {
		panic(err)
	}

	err = db.InitDB()
	if err != nil {
		panic(err)
	}

	r := gin.Default()

	routes.Init(r)
	err = r.Run(fmt.Sprintf(":%d", cfg.Server.Port))
	if err != nil {
		panic(err)
	}

}
