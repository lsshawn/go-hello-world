package main

import (
	"context"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/lsshawn/go-todo/templates"
)

func main() {
	app := echo.New()

	component := templates.Index()
	component.Render(context.Background(), os.Stdout)

	app.GET("/", func(c echo.Context) error {
		return component.Render(context.Background(), c.Response().Writer)
	})

	app.Static("/css", "css")
	app.Static("/static", "static")

	app.Logger.Fatal(app.Start(":1323"))
}
