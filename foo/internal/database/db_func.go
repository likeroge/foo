package database

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func InitDB() error {
	db, err := sql.Open("./sqlite3", "foo.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	result, err := db.Exec("CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY AUTOINCREMENT, name TEXT, email TEXT)")
	if err != nil {
		panic(err)
	}
	fmt.Printf("result: %d\n", result)
	return nil
}

func GetDB() *sql.DB {
	db, err := sql.Open("sqlite3", "foo.db")

	if err != nil {
		panic(err)
	}
	return db
}
