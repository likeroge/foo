package utils

import (
	"ego.dev21/greetings/internal/entities"
)

func GetUsersListSize(users []entities.User) int {
	l := len(users)
	return l
}