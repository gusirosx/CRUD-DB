package router

import (
	"clean2/adapters"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func NewRouter(e *echo.Echo, c adapters.AppController) *echo.Echo {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/users", func(context echo.Context) error { return c.User.GetUsers(context) })

	return e
}
