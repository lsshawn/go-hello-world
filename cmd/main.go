package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/lsshawn/go-todo/dto"
	"github.com/lsshawn/go-todo/handler"
	"github.com/lsshawn/go-todo/view"
	"github.com/lsshawn/go-todo/view/components"
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

	app.POST("/add-todo", func(c echo.Context) error {
		text := c.FormValue("text")
		fmt.Printf("LS -> cmd/main.go:41 -> text: %+v\n", text)

		// mimic a server call
		time.Sleep(1 * time.Second)
		newTodo := &dto.TodoCardDto{
			Text: text,
		}

		component := components.TodoCard(newTodo)
		return component.Render(context.Background(), c.Response().Writer)
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
