package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"ego.dev21/greetings/internal/entities"
	"ego.dev21/greetings/internal/repository"
)

type UserHandler struct {
	Repository *repository.Repositories
}

func NewUserHandler(repos *repository.Repositories) *UserHandler {
	return &UserHandler{
		Repository: repos,
	}
}

func (h *UserHandler) GetHello(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	allUsers := h.Repository.UserRepository.GetAllUsers()
	resp, err := json.Marshal(allUsers)
	if err != nil {
		log.Fatal(err)
	}
	w.Write(resp)
}

func GetHello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}

func (h *UserHandler) FindAllUsers(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		users := h.Repository.UserRepository.GetAllUsers()
		// repository := sqliterepos.NewUserSqliteRepository()
		// users := repository.GetAllUsers()
		json.NewEncoder(w).Encode(users)
	}
}

func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
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
		h.Repository.UserRepository.DeleteUser(intVal)
		// repository := sqliterepos.NewUserSqliteRepository()
		// repository.DeleteUser(intVal)
		w.Write([]byte("User deleted!"))
	}
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
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
		// repository := sqliterepos.NewUserSqliteRepository()

		// result, err := repository.AddUser(entities.User{Name: u.Name, Email: u.Email})
		result, err := h.Repository.UserRepository.AddUser(entities.User{Name: u.Name, Email: u.Email})
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(result)
		w.Write([]byte("User added!"))
	}

}

func (h *UserHandler) FindUserById(w http.ResponseWriter, r *http.Request) {
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
		// repository := sqliterepos.NewUserSqliteRepository()
		// user, err := repository.FindUserById(intVal)
		user, err := h.Repository.UserRepository.FindUserById(intVal)
		if err != nil {
			fmt.Println(err)
		}
		json.NewEncoder(w).Encode(user)
	}
}

func (h *UserHandler) FindUserByName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	if r.Method == "GET" {
		pathVal := r.PathValue("userName")
		fmt.Println(pathVal)
		// repository := sqliterepos.NewUserSqliteRepository()
		// user, err := repository.FindUserByName(pathVal)
		user, err := h.Repository.UserRepository.FindUserByName(pathVal)
		if err != nil {
			fmt.Println(err)
		}
		json.NewEncoder(w).Encode(user)
	}
}

func (h *UserHandler) FindUserByEmail(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	if r.Method == "GET" {
		pathVal := r.PathValue("userEmail")
		fmt.Println(pathVal)
		// repository := sqliterepos.NewUserSqliteRepository()
		// user, err := repository.FindUserByEmail(pathVal)
		user, err := h.Repository.UserRepository.FindUserByEmail(pathVal)
		if err != nil {
			fmt.Println(err)
		}
		json.NewEncoder(w).Encode(user)
	}
}
