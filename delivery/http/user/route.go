package user

import (
	"github.com/labstack/echo/v4"
	"go-base/usecase"
)

type Route struct {
	useCase *usecase.UseCase
}

func Init(group *echo.Group, useCase *usecase.UseCase) {
	r := &Route{useCase: useCase}

	group.POST("/create", r.Create)
	group.POST("/login", r.Login)
	group.POST("/refresh", r.RefreshToken)
}
