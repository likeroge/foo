package entities

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Id    int    `json:"id"`
}