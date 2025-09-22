package internal

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"ego.dev21/greetings/internal/entities"
	"ego.dev21/greetings/internal/repository"
	"ego.dev21/greetings/internal/utils"
)

func GetHello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}

func FindAllUsers(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET"{
		//all users
		w.Header().Set("Content-Type", "application/json")
		users := repository.GetAllUsers()
		json.NewEncoder(w).Encode(users)
	}
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		buf := make([]byte, r.ContentLength)
		log.Println("POST")
		r.Body.Read(buf)
		fmt.Println(buf)
		fmt.Println(string(buf))
		var u entities.User
		json.Unmarshal(buf, &u)
		fmt.Println(u)
		fmt.Println(u.Name, u.Email)
		users := repository.GetAllUsers()
		if utils.GetUsersListSize(users) >= 10 {
			w.Write([]byte("Too many users!"))
			fmt.Println("Too many users!")
			return
		}
		repository.AddUser(entities.User{Name: u.Name, Email: u.Email})
	} else{
		w.Write([]byte("Page not found!"))
	}
}