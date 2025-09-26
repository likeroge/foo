package users

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"ego.dev21/greetings/internal/entities"
	sqliterepos "ego.dev21/greetings/internal/repository/sqlite_repos"
)

func GetHello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}

func FindAllUsers(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		repository := sqliterepos.NewUserSqliteRepository()
		users := repository.GetAllUsers()
		json.NewEncoder(w).Encode(users)
	}
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	if r.Method == "DELETE" {
		pathVal := r.PathValue("userId")
		intVal, err := strconv.Atoi(pathVal)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(pathVal)
		repository := sqliterepos.NewUserSqliteRepository()
		repository.DeleteUser(intVal)
		w.Write([]byte("User deleted!"))
	}
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "*")
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
		repository := sqliterepos.NewUserSqliteRepository()

		result, err := repository.AddUser(entities.User{Name: u.Name, Email: u.Email})
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(result)
		w.Write([]byte("User added!"))
	}

}

func FindUserById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	if r.Method == "GET" {
		pathVal := r.PathValue("userId")
		intVal, err := strconv.Atoi(pathVal)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(pathVal)
		repository := sqliterepos.NewUserSqliteRepository()
		user, err := repository.FindUserById(intVal)
		if err != nil {
			fmt.Println(err)
		}
		json.NewEncoder(w).Encode(user)
	}
}

func FindUserByName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	if r.Method == "GET" {
		pathVal := r.PathValue("userName")
		fmt.Println(pathVal)
		repository := sqliterepos.NewUserSqliteRepository()
		user, err := repository.FindUserByName(pathVal)
		if err != nil {
			fmt.Println(err)
		}
		json.NewEncoder(w).Encode(user)
	}
}

func FindUserByEmail(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	if r.Method == "GET" {
		pathVal := r.PathValue("userEmail")
		fmt.Println(pathVal)
		repository := sqliterepos.NewUserSqliteRepository()
		user, err := repository.FindUserByEmail(pathVal)
		if err != nil {
			fmt.Println(err)
		}
		json.NewEncoder(w).Encode(user)
	}
}
