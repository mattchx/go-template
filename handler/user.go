package handler

import (
	"go-project-template/model"
	"go-project-template/view/user"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
}

func (h UserHandler) HandlerUserShow(c echo.Context) error {
	userView := model.User{
		Email: "matt@ch.com",
}

	return render(c, user.Show(userView))
}
