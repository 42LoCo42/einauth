package server

import (
	"github.com/labstack/echo/v4"
)

func Init() (*echo.Echo, error) {
	e := echo.New()

	e.GET("/", UI)
	e.GET("/auth", Auth)
	e.POST("/login", Login)

	return e, nil
}
