package repository

import "ego.dev21/greetings/internal/entities"

var users = []entities.User{}

func GetAllUsers() []entities.User {
	return users
}

func AddUser(user entities.User) {
	users = append(users, user)
}