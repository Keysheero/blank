package main

import (
	application "gostart/internal/app"
)

func main() {
	app := application.InitializeApplication()
	app.Run()
}
