package main

import (
	"context"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/lsshawn/go-todo/dto"
	"github.com/lsshawn/go-todo/handler"
	"github.com/lsshawn/go-todo/view"
)

func main() {
	todos := []*dto.TodoCardDto{
		{
			ID:      uuid.New().String(),
			Text:    "First item",
			Checked: false,
		}, {
			ID:      uuid.New().String(),
			Text:    "Second item",
			Checked: false,
		},
	}

	app := echo.New()

	// component := view.Index()
	// component.Render(context.Background(), os.Stdout)

	userHandler := handler.UserHandler{}

	app.Use(userMiddleware)
	app.GET("/user", userHandler.HandleUserShow)

	app.GET("/ping", func(c echo.Context) error {
		return c.String(200, "pong")
	})

	app.GET("/", func(c echo.Context) error {
		component := view.Index(todos)
		return component.Render(context.Background(), c.Response().Writer)
	})

	app.Static("/css", "css")
	app.Static("/static", "static")

	app.Logger.Fatal(app.Start(":1323"))
}

func userMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := context.WithValue(c.Request().Context(), "user", "a@gg.com")
		c.SetRequest(c.Request().WithContext(ctx))
		return next(c)
	}
}
