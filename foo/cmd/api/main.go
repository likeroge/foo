package main

import (
	"ego.dev21/greetings/internal/app"
)

func main() {
	app := app.NewApplication()
	defer app.Db.Close()
	app.Run()
}
