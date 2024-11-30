package services

import "web4/internal/models"

var users map[int]models.User = make(map[int]models.User)
var id int = 1

func AddUser(user models.User) {
	user.ID = id
	users[id] = user
	id++
}

func GetUsers() []models.User {
	var userList []models.User
	for _, user := range users {
		userList = append(userList, user)
	}
	return userList
}
