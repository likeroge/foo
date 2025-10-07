package repository

import (
	"database/sql"

	userrepos "ego.dev21/greetings/internal/repository/user_repos"
)

type Repositories struct {
	UserRepository *userrepos.UserSqliteRepository
}

func NewRepositories(db *sql.DB) *Repositories {
	return &Repositories{
		UserRepository: userrepos.NewUserSqliteRepository(db),
	}
}
