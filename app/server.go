package app

import (
	"github.com/labstack/echo"
)

func Create(controller Controller) *echo.Echo {
	e := echo.New()
	e.GET("/hello", controller.Get)
	return e
}

type Controller interface {
	Get(c echo.Context) error
}
