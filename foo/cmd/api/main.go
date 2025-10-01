package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"ego.dev21/greetings/internal/database"
	"ego.dev21/greetings/internal/presentation/http/handlers/files"
	"ego.dev21/greetings/internal/presentation/http/handlers/ofp"
	"ego.dev21/greetings/internal/presentation/http/handlers/users"
)

//	func main() {
//		ofpParser.ParseOfp("Hello world asdasdas asdasd asdasd , 123123")
//	}
func main() {
	http.HandleFunc("/hello", users.GetHello)
	// users
	http.HandleFunc("/api/user/all", users.FindAllUsers)
	http.HandleFunc("/api/user", users.CreateUser)
	http.HandleFunc("/api/user/delete/{userId}", users.DeleteUser)
	http.HandleFunc("/api/user/find/name/{userName}", users.FindUserByName)
	http.HandleFunc("/api/user/find/email/{userEmail}", users.FindUserByEmail)
	http.HandleFunc("/api/user/find/id/{userId}", users.FindUserById)

	//files
	http.HandleFunc("/api/file/send", files.SendFile)

	//ofp
	http.HandleFunc("/api/ofp/send", ofp.PostOfpToBackend)

	err := database.InitDB()
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}
	log.Println("Server started on http://localhost:5000")
	er := http.ListenAndServe(":5000", nil)
	if er != nil {
		panic(er)
	}
}
