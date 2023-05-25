package user

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"go-base/model"
	"go-base/util"
	"net/http"
)

func (r *Route) Login(c echo.Context) error {
	var (
		req model.LoginReq
		err error
		res *util.Response
	)

	if err = c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, fmt.Sprintf("Error when parse req: %s", err.Error()))
	}

	//check valid
	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	//login
	res, err = r.useCase.UserUseCase.Login(c, req)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, res)
}
