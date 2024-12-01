package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

var DB *sql.DB

func InitDB() error {
	log.Printf("Starting DB initialization")
	var err error
	DB, err = sql.Open("postgres", "user=alex password=0000 dbname=postgres sslmode=disable")
	if err != nil {
		return err
	}

	return DB.Ping()
}
