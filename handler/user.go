package handler

import (
	"go-todo/view/user"

	"github.com/labstack/echo/v4"
)

type UserHandler struct{}

func (h UserHandler) HandleUserShow(c echo.Context) error {
	return user.Show().Render(c.Request().Context(), c.Response())
	return nil
}
