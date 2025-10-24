package app

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"ego.dev21/greetings/internal/database"
	"ego.dev21/greetings/internal/handler"
	"ego.dev21/greetings/internal/repository"
)

type Application struct {
	MainHandler *http.Handler
	// Server       *http.Server
	Db           *sql.DB
	Port         string
	Repositories *repository.Repositories
	LogFile      *os.File
}

func NewApplication() *Application {
	app := &Application{}

	// server := &http.Server{}
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}
	app.SetupDatabase()
	app.SetupRepositories()

	router := handler.SetupRoutes(app.Repositories)

	app.Port = port
	app.MainHandler = router.HttpHandler
	// app.Server = server

	return app
}

func (app *Application) SetupLogToFile() {
	file, err := os.OpenFile("logs/log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(file)
	log.Println("Log file created")
	app.LogFile = file
}

func (app *Application) SetupDatabase() {
	// err := database.InitDB()
	_, err := database.ExecuteSQLFileLineByLine("./migrations/init_migrations.sql")
	if err != nil {
		log.Println("Error executing statement:", err)
		log.Fatal(err)
	}

	app.Db = database.GetDB()
	// app.Db = db
}

func (app *Application) SetupRepositories() {
	app.Repositories = repository.NewRepositories(app.Db)
}

// func (app *Application) SetupHandlers() {

// }

func (app *Application) Run() {
	log.Println("Server started on http://localhost" + ":" + app.Port)
	err := http.ListenAndServe(":"+app.Port, *app.MainHandler)
	if err != nil {
		log.Fatal(err)
	}
}
