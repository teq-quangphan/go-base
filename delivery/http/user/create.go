package user

import (
	"github.com/labstack/echo/v4"
	"go-base/model"
)

func (r *Route) Create(c echo.Context) error {
	if c != nil {
		var user model.User
		_ = r.useCase.UserUseCase.Create(user)
	}
	return nil
}
