package server

import (
	"net/http"

	"github.com/42LoCo42/einauth/utils"
	"github.com/labstack/echo/v4"
)

func Logout(c echo.Context) error {
	c.SetCookie(utils.ClearCookie("token"))
	return c.Redirect(http.StatusSeeOther, "/")
}
