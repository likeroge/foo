package internal

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

var users = []User{}

func GetAllUsers() []User {
	return users
}

func AddUser(user User) {
	users = append(users, user)
}