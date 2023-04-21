package user

import (
	"fmt"
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

	////check valid
	//if err = c.Validate(&req); err != nil {
	//	return c.JSON(http.StatusBadRequest, fmt.Sprintf("Error when validate req: %s", err.Error()))
	//}

	//login
	res, err = r.useCase.UserUseCase.Login(c, req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, res)
}
