package http

import (
	"net/http"
	"regexp"

	userHandler "go-base/delivery/http/user"
	"go-base/usecase"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go-base/delivery/http/healthcheck"
)

func NewHTTPHandler(useCase *usecase.UseCase) *echo.Echo {
	var (
		e = echo.New()
	)

	e.Use(middleware.Recover())
	e.Pre(middleware.RemoveTrailingSlash())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		Skipper: middleware.DefaultSkipper,
		AllowOriginFunc: func(origin string) (bool, error) {
			return regexp.MatchString(
				`^https:\/\/(|[a-zA-Z0-9]+[a-zA-Z0-9-._]*[a-zA-Z0-9]+\.)teqnological.asia$`,
				origin,
			)
		},
		AllowMethods: []string{
			http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch,
			http.MethodPost, http.MethodDelete, http.MethodOptions,
		},
	}))

	// Health check use for microservice
	healthcheck.Init(e.Group("/health-check"))

	// API docs
	//if !config.GetConfig().Stage.IsProd() {
	//	e.GET("/docs/*", echoSwagger.WrapHandler)
	//}

	// APIs
	api := e.Group("/api")
	userHandler.Init(api.Group("/user"), useCase)

	return e
}
