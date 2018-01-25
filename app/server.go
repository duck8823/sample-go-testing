package app

import (
	"github.com/labstack/echo"
	"net/http"
	"fmt"
)

func CreateServer() *echo.Echo {
	e := echo.New()
	e.GET("/hello", hello)
	return e
}

func hello(c echo.Context) error {
	message := fmt.Sprintf("Hello %s.", c.QueryParam("name"))
	return c.String(http.StatusOK, message)
}