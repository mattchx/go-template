package main

import (
	"github.com/labstack/echo/v4"
	// "templ-practice/view/layout"
	"go-project-template/handler"
)

// type DB struct {}

func main() {
	app := echo.New()

	// var db DB
	// userHandler := handler.UserHandler{ db }
	userHandler := handler.UserHandler{}
	app.GET("/", userHandler.HandlerUserShow)

	app.Start(":8080")
	// component := layout.Base()
	// component.Render(context.Background(), os.Stdout)
}
