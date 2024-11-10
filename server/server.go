package server

import "github.com/labstack/echo/v4"

func Init() (e *echo.Echo, err error) {
	e = echo.New()

	return e, nil
}
