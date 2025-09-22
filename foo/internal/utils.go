package internal

func GetUsersListSize() int {
	users := GetAllUsers()
	return len(users)
}