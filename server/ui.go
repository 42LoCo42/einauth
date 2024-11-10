package server

import (
	"log"
	"net/http"

	"github.com/42LoCo42/einauth/utils"
	"github.com/labstack/echo/v4"
)

func UI(c echo.Context) error {
	cookie, err := c.Cookie("einauth")
	if err != nil {
		return LoginUI(c)
	}

	user, err := utils.VerifyCookie[CookieUser](cookie)
	if err != nil {
		return LoginUI(c)
	}

	log.Print(user)

	return c.String(http.StatusOK, "TODO main page")
}

func LoginUI(c echo.Context) error {
	return c.String(http.StatusForbidden, "TODO login page")
}
