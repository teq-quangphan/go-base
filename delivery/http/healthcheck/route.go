package healthcheck

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func Init(group *echo.Group) {
	group.GET("", func(ctx echo.Context) error {
		return ctx.NoContent(http.StatusOK)
	})
}
