package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"web4/internal/config"
)

var DB *sql.DB

func InitDB() error {
	log.Printf("Starting DB initialization")
	cfg := config.GetConfig()
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s", cfg.Database.User, cfg.Database.Password, cfg.Database.DBName, cfg.Database.SSLMode)
	
	var err error
	DB, err = sql.Open("postgres", dsn)
	if err != nil {
		return err
	}

	return DB.Ping()
}
