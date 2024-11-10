package server

import (
	"github.com/labstack/echo/v4"
)

func Init() (*echo.Echo, error) {
	e := echo.New()

	e.GET("/", UI)
	e.POST("/login", Login)
	// TODO auth endpoint

	return e, nil
}
