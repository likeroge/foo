package handler

import (
	"net/http"

	"ego.dev21/greetings/internal/handler/middleware"
	"ego.dev21/greetings/internal/repository"
)

type Router struct {
	UserHandler  *UserHandler
	OFPHandler   *OfpHandler
	FilesHandler *FilesHandler

	HttpHandler  http.Handler
	Repositories *repository.Repositories
}

type Middleware func(http.Handler) http.Handler

// Chain: mws[0] выполняется первым
func Chain(h http.Handler, mws ...Middleware) *http.Handler {
	for i := len(mws) - 1; i >= 0; i-- {
		h = mws[i](h)
	}
	return &h
}

// func SetupRoutes(repositories *repository.Repositories) *Router {
func SetupRoutes(repositories *repository.Repositories) *Router {

	r := &Router{}

	r.Repositories = repositories
	r.UserHandler = NewUserHandler(r.Repositories)
	r.OFPHandler = NewOfpHandler(r.Repositories)
	r.FilesHandler = NewFilesHandler(r.Repositories)

	topServerMux := http.NewServeMux()
	protectedServerMux := http.NewServeMux()
	publicServerMux := http.NewServeMux()

	////////////
	publicServerMux.HandleFunc("/hello", r.UserHandler.GetHello)

	//user
	protectedServerMux.HandleFunc("/user/all", r.UserHandler.FindAllUsers)
	protectedServerMux.HandleFunc("/user", r.UserHandler.CreateUser)
	protectedServerMux.HandleFunc("/user/delete/{userId}", r.UserHandler.DeleteUser)
	protectedServerMux.HandleFunc("/user/find/name/{userName}", r.UserHandler.FindUserByName)
	protectedServerMux.HandleFunc("/user/find/email/{userEmail}", r.UserHandler.FindUserByEmail)
	protectedServerMux.HandleFunc("/user/find/id/{userId}", r.UserHandler.FindUserById)

	//ofp
	protectedServerMux.HandleFunc("/ofp/all", r.OFPHandler.GetAllOFPInfo)
	protectedServerMux.HandleFunc("/ofp/{id}", r.OFPHandler.GetOFPInfoById)
	protectedServerMux.HandleFunc("/ofp/send", r.OFPHandler.PostOfpToBackend)
	protectedServerMux.HandleFunc("/ofp/delete/{id}", r.OFPHandler.DeleteOFPInfoById)

	// file
	protectedServerMux.HandleFunc("/file/send", r.OFPHandler.PostOfpToBackend)

	// configure server-mux routes
	topServerMux.Handle("/", publicServerMux)
	// topServerMux.Handle("/api/", http.StripPrefix("/api", *Chain(protectedServerMux, middleware.AuthMiddleware)))
	topServerMux.Handle("/api/", http.StripPrefix("/api", protectedServerMux))

	r.HttpHandler = *Chain(topServerMux, middleware.LoggerMiddleware)

	return r
}
