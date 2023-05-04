package user

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"go-base/model"
	"net/http"
)

func (r *Route) Create(c echo.Context) error {
	user := &model.User{}

	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	newUser, err := r.useCase.UserUseCase.Create(*user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, newUser)
}
