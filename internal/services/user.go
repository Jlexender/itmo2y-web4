package services

import (
	"github.com/Masterminds/squirrel"
	"log"
	"web4/internal/db"
	"web4/internal/models"
)

func AddUser(user models.User) {
	query := squirrel.Insert("users").Columns("name", "hash").Values(user.Name, user.Hash).PlaceholderFormat(squirrel.Dollar)

	_, err := query.RunWith(db.DB).Exec()
	if err != nil {
		log.Printf("Error adding user: %v", err)
	}
}

func UserExists(name string) bool {
	query := squirrel.Select("name").From("users").Where(squirrel.Eq{"name": name}).PlaceholderFormat(squirrel.Dollar)

	rows, err := query.RunWith(db.DB).Query()
	if err != nil {
		log.Printf("Error checking user: %v", err)
		return false
	}

	defer rows.Close()

	return rows.Next()
}

func GetUser(name string) (bool, models.User) {
	query := squirrel.Select("name", "hash").From("users").Where(squirrel.Eq{"name": name}).PlaceholderFormat(squirrel.Dollar)

	row := query.RunWith(db.DB).QueryRow()

	var u models.User
	err := row.Scan(&u.Name, &u.Hash)
	if err != nil {
		log.Printf("Error getting user: %v", err)
		return false, u
	}

	return true, u
}
