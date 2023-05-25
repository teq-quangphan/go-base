package user

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func (r *Route) GetOneUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	newUser, err := r.useCase.UserUseCase.GetOneUser(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, newUser)
}
