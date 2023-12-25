package handler

import (
	"github.com/labstack/echo/v4"
)

func render(c echo.Context, component templ.component) error {
	return component.Render(c.Request().Context(), c.Response())
}
