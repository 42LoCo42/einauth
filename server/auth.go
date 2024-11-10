package server

import (
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/42LoCo42/einauth/config"
	"github.com/42LoCo42/einauth/utils"
	"github.com/labstack/echo/v4"
)

func Auth(c echo.Context) error {
	headers := c.Request().Header
	target := url.URL{
		Scheme: headers.Get("x-forwarded-proto"),
		Host:   headers.Get("x-forwarded-host"),
		Path:   headers.Get("x-forwarded-uri"),
	}

	redir := func() error {

		return c.Redirect(
			http.StatusSeeOther,
			fmt.Sprintf("%s?target=%s",
				config.CONFIG.URL,
				url.QueryEscape(target.String()),
			),
		)
	}

	cookie, err := c.Cookie("einauth")
	if err != nil {
		return redir()
	}

	user, err := utils.VerifyCookie[CookieUser](cookie)
	if err != nil {
		return redir()
	}

	// TODO verify that user has access
	log.Print(user)
	return nil
}
