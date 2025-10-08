package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	_ "modernc.org/sqlite"
)

func InitDB() error {
	db, err := sql.Open("sqlite", "./data/database.db")
	if err != nil {
		log.Println(err)
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY AUTOINCREMENT, name TEXT, email TEXT)")
	if err != nil {
		log.Println(err)
		panic(err)
	}
	return nil
}

func ExecuteSQLFileLineByLine(filePath string) (*sql.DB, error) {
	db, err := sql.Open("sqlite", "./data/database.db")
	if err != nil {
		log.Println(err)
		panic(err)
	}
	defer db.Close()

	sqlScript, err := os.ReadFile(filePath)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	// Разбиваем скрипт на отдельные команды
	statements := strings.Split(string(sqlScript), ";")

	for _, stmt := range statements {
		stmt = strings.TrimSpace(stmt)
		if stmt == "" {
			continue
		}
		fmt.Println(stmt)
		_, err := db.Exec(stmt)
		if err != nil {
			log.Println(err)
			return nil, err
		}
	}
	return db, nil
}

func GetDB() *sql.DB {
	db, err := sql.Open("sqlite", "./data/database.db")
	if err != nil {
		panic(err)
	}
	return db
}
