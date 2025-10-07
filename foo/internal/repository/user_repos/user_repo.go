package userrepos

import "ego.dev21/greetings/internal/entities"

type UserRepository interface {
	GetAllUsers() []entities.User
	AddUser(user entities.User) (int64, error)
	DeleteUser(id int)
	FindUserById(id int) (*entities.User, error)
	FindUserByName(name string) (*entities.User, error)
	FindUserByEmail(email string) (*entities.User, error)
}
