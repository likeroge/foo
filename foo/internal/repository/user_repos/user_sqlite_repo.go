package userrepos

import (
	"database/sql"
	"fmt"
	"log"

	"ego.dev21/greetings/internal/entities"
)

type UserSqliteRepository struct {
	db *sql.DB
}

func NewUserSqliteRepository(db *sql.DB) *UserSqliteRepository {
	return &UserSqliteRepository{
		db: db,
	}
}

func (r *UserSqliteRepository) DeleteUser(id int) {
	result, err := r.db.Exec("DELETE FROM users WHERE id = ?", id)
	if err != nil {
		panic(err)
	}
	fmt.Printf("result: %d\n", result)
}

func (r *UserSqliteRepository) GetAllUsers() []entities.User {
	// db := database.GetDB()
	// defer db.Close()

	rows, err := r.db.Query("SELECT id, name, email FROM users")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var users []entities.User = []entities.User{}
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
	sqlStr := `
					INSERT INTO users (name, email, created_at, updated_at)
					VALUES (?, ?, datetime('now'), datetime('now'))
					ON CONFLICT(email) 
					DO UPDATE SET
						name = excluded.name,
						updated_at = datetime('now');
				`
	result, err := r.db.Exec(sqlStr, user.Name, user.Email)
	if err != nil {
		log.Println("Error due to insert user", err)
		return -1, err
	}
	lastInsertedId, _ := result.LastInsertId()
	return lastInsertedId, nil
}

func (r *UserSqliteRepository) FindUserById(id int) (*entities.User, error) {
	rows := r.db.QueryRow("SELECT id, name, email FROM users WHERE id = ?", id)
	var user entities.User
	if err := rows.Scan(&user.Id, &user.Name, &user.Email); err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserSqliteRepository) FindUserByName(name string) (*entities.User, error) {
	rows := r.db.QueryRow("SELECT id, name, email FROM users WHERE name = ?", name)
	var user entities.User
	if err := rows.Scan(&user.Id, &user.Name, &user.Email); err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserSqliteRepository) FindUserByEmail(email string) (*entities.User, error) {
	rows := r.db.QueryRow("SELECT id, name, email FROM users WHERE email = ?", email)
	var user entities.User
	if err := rows.Scan(&user.Id, &user.Name, &user.Email); err != nil {
		return nil, err
	}
	return &user, nil
}
