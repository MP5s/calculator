package main

import (
	app "github.com/MP5s/calculator/internal/application"
)

func main() {
	a := app.New()
	a.RunServer() // Запускаем приложение
}
