package main

import (
	"fmt"

	"github.com/labstack/echo/v4"
	// "templ-practice/view/layout"
	"go-project-template/handler"
	"go-project-template/view/component"
)

// type DB struct {}

func main() {
	app := echo.New()

	// var db DB
	// userHandler := handler.UserHandler{ db }
	userHandler := handler.UserHandler{}
	app.GET("/", userHandler.HandlerUserShow)

	app.POST("/clicked", func(c echo.Context) error {
		fmt.Println("*** clicked ***")
		return component.Button("HTMX BUTTON").Render(c.Request().Context(), c.Response())
	})

	app.Start(":8080")
	// component := layout.Base()
	// component.Render(context.Background(), os.Stdout)
}
