package sqliterepos

import (
	"database/sql"
	"fmt"

	"ego.dev21/greetings/internal/entities"
	"ego.dev21/greetings/internal/repository"
)

type UserSqliteRepository struct {
	// db *sql.DB
}

func NewUserSqliteRepository() repository.UserRepository {
	return &UserSqliteRepository{}
}

func (r *UserSqliteRepository) DeleteUser(id int) {
	db, err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	result, err := db.Exec("DELETE FROM users WHERE id = ?", id)
	if err != nil {
		panic(err)
	}
	fmt.Printf("result: %d\n", result)
}

func (r *UserSqliteRepository) GetAllUsers() []entities.User {
	db, err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT id, name, email FROM users")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var users []entities.User
	for rows.Next() {
		var user entities.User
		if err := rows.Scan(&user.Id, &user.Name, &user.Email); err != nil {
			panic(err)
		}
		users = append(users, user)
	}
	return users
}

func (r *UserSqliteRepository) AddUser(user entities.User) (int64, error) {
	db, err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
		return -1, err
	}
	defer db.Close()

	result, err := db.Exec("INSERT INTO users (name, email) VALUES (?, ?)", user.Name, user.Email)
	if err != nil {
		return -1, err
	}

	return result.LastInsertId()
}

func (r *UserSqliteRepository) FindUserById(id int) (*entities.User, error) {
	db, err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows := db.QueryRow("SELECT id, name, email FROM users WHERE id = ?", id)
	var user entities.User
	if err := rows.Scan(&user.Id, &user.Name, &user.Email); err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserSqliteRepository) FindUserByName(name string) (*entities.User, error) {
	db, err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows := db.QueryRow("SELECT id, name, email FROM users WHERE name = ?", name)
	var user entities.User
	if err := rows.Scan(&user.Id, &user.Name, &user.Email); err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserSqliteRepository) FindUserByEmail(email string) (*entities.User, error) {
	db, err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows := db.QueryRow("SELECT id, name, email FROM users WHERE email = ?", email)
	var user entities.User
	if err := rows.Scan(&user.Id, &user.Name, &user.Email); err != nil {
		return nil, err
	}
	return &user, nil
}
