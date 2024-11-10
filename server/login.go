package server

import (
	"net/http"
	"reflect"

	"github.com/42LoCo42/einauth/db"
	"github.com/42LoCo42/einauth/utils"
	"github.com/go-faster/errors"
	"github.com/labstack/echo/v4"
)

type CookieUser struct {
	ID      uint
	Name    string
	IsAdmin bool
	Groups  []string
}

func Login(c echo.Context) error {
	reject := func() error {
		return echo.NewHTTPError(
			http.StatusForbidden,
			"invalid credentials",
		)
	}

	username := c.FormValue("username")
	password := c.FormValue("password")

	var user db.User
	if err := db.DB.
		Preload(reflect.TypeOf(db.Group{}).Name()+"s").
		First(&user, &db.User{Name: username}).
		Error; err != nil {
		return reject()
	}

	if ok, err := utils.VerifyPassword(password, user.Password); err != nil || !ok {
		return reject()
	}

	cookie, err := utils.SignCookie(CookieUser{
		ID:      user.ID,
		Name:    user.Name,
		IsAdmin: user.IsAdmin,
		Groups: utils.Map(user.Groups, func(group db.Group) string {
			return group.Name
		}),
	})
	if err != nil {
		return echo.NewHTTPError(
			http.StatusInternalServerError,
			errors.Wrap(err, "could not create cookie"),
		)
	}

	c.SetCookie(cookie)

	redir, err := c.Cookie("einauth-redir")
	if err != nil {
		// TODO better error indication
		return c.Redirect(http.StatusSeeOther, "/")
	}

	c.SetCookie(utils.ClearCookie("redir"))
	return c.Redirect(http.StatusSeeOther, redir.Value)
}
