package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "modernc.org/sqlite"
)

func InitDB() error {
	db, err := sql.Open("sqlite", "./foo.db")
	if err != nil {
		log.Println(err)
		panic(err)
	}
	defer db.Close()
	result, err := db.Exec("CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY AUTOINCREMENT, name TEXT, email TEXT)")
	if err != nil {
		log.Println(err)
		panic(err)
	}
	fmt.Printf("result: %d\n", result)
	return nil
}

func GetDB() *sql.DB {
	db, err := sql.Open("sqlite", "./foo.db")
	if err != nil {
		panic(err)
	}
	return db
}
