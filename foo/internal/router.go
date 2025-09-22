package internal

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func GetHello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}

func FindAllUsers(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET"{
		//all users
		w.Header().Set("Content-Type", "application/json")
		users := GetAllUsers()
		for _, user := range users {
			userBytes, _ := json.Marshal(user)
			fmt.Fprintln(w, string(userBytes))
		}
	}
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		buf := make([]byte, r.ContentLength)
		log.Println("POST")
		r.Body.Read(buf)
		fmt.Println(buf)
		fmt.Println(string(buf))
		var u User
		json.Unmarshal(buf, &u)
		fmt.Println(u)
		fmt.Println(u.Name, u.Email)
		if GetUsersListSize() >= 10 {
			w.Write([]byte("Too many users!"))
			fmt.Println("Too many users!")
			return
		}
		AddUser(User{Name: u.Name, Email: u.Email})
	} else{
		w.Write([]byte("Page not found!"))
	}
}