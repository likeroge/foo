package main

import (
	"log"
	"net/http"

	"ego.dev21/greetings/internal"
)



func main() {
	http.HandleFunc("/hello", internal.GetHello)
	http.HandleFunc("/users", internal.FindAllUsers)
	http.HandleFunc("/user", internal.CreateUser)

	log.Println("Server started on http://localhost:5000")
	http.ListenAndServe(":5000", nil)
}