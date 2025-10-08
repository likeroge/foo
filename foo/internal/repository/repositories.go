package repository

import (
	"database/sql"

	ofprepos "ego.dev21/greetings/internal/repository/ofp_repos"
	userrepos "ego.dev21/greetings/internal/repository/user_repos"
)

type Repositories struct {
	UserRepository *userrepos.UserSqliteRepository
	OFPRpository   *ofprepos.OFPSQLiteRepo
}

func NewRepositories(db *sql.DB) *Repositories {
	return &Repositories{
		UserRepository: userrepos.NewUserSqliteRepository(db),
		OFPRpository:   ofprepos.NewOFPSQLiteRepo(db),
	}
}
