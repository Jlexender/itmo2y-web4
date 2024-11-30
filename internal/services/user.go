package services

import "web4/internal/models"

var users map[string]models.User = make(map[string]models.User)

func AddUser(user models.User) {
	users[user.User] = user
}

func UserExists(user string) bool {
	_, ok := users[user]
	return ok
}

func GetUser(user string) models.User {
	return users[user]
}
