package main

import (
	"ego.dev21/greetings/internal/app"
)

func main() {
	app := app.NewApplication()
	app.SetupLogToFile()
	defer app.Db.Close()
	defer app.LogFile.Close()
	app.Run()
}
