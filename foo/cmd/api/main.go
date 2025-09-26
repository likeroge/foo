package main

import (
	"log"
	"net/http"

	"ego.dev21/greetings/internal/database"
	"ego.dev21/greetings/internal/presentation/http/handlers/users"
)

// type SqliteDB struct {
// 	db *sql.DB
// }

func main() {
	http.HandleFunc("/hello", users.GetHello)
	//users
	http.HandleFunc("/api/user/all", users.FindAllUsers)
	http.HandleFunc("/api/user", users.CreateUser)
	http.HandleFunc("/api/user/delete/{userId}", users.DeleteUser)
	http.HandleFunc("/api/user/find/name/{userName}", users.FindUserByName)
	http.HandleFunc("/api/user/find/email/{userEmail}", users.FindUserByEmail)
	http.HandleFunc("/api/user/find/id/{userId}", users.FindUserById)

	err := database.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Server started on http://localhost:5000")
	http.ListenAndServe(":5000", nil)
}
