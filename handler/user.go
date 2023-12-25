package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/lsshawn/go-todo/model"
	"github.com/lsshawn/go-todo/view/user"
)

type UserHandler struct{}

func (h UserHandler) HandleUserShow(c echo.Context) error {
	u := model.User{
		Email: "l@sshawn.com",
	}
	return render(c, user.Show(u))
}
