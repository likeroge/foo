package main

import (
	"log"
	"net/http"
	"os"

	"ego.dev21/greetings/internal/database"
	"ego.dev21/greetings/internal/presentation/http/handlers/users"
)

func main() {
	http.HandleFunc("/hello", users.GetHello)
	// users
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
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}
	log.Println("Server started on http://localhost:5000")
	http.ListenAndServe(":"+port, nil)
}
