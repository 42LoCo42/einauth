package server

import (
	"encoding/json"
	"log"
	"net/http"
	"reflect"

	"github.com/42LoCo42/einauth/db"
	"github.com/42LoCo42/einauth/utils"
	"github.com/go-faster/errors"
	"github.com/labstack/echo/v4"
)

type LoginReq struct {
	Username string
	Password string
}

type CookieUser struct {
	ID      uint
	Name    string
	IsAdmin bool
	Groups  []string
}

func Login(c echo.Context) error {
	credsErr := func() error {
		return echo.NewHTTPError(
			http.StatusForbidden,
			"invalid credentials",
		)
	}

	internalErr := func(err error, msg string) error {
		return echo.NewHTTPError(
			http.StatusInternalServerError,
			errors.Wrap(err, msg),
		)
	}

	var req LoginReq
	if err := json.
		NewDecoder(c.Request().Body).
		Decode(&req); err != nil {
		log.Print("1", err)
		return credsErr()
	}

	if req.Username == "" || req.Password == "" {
		log.Print("2")
		return credsErr()
	}

	var user db.User
	if err := db.DB.
		Preload(reflect.TypeOf(db.Group{}).Name()+"s").
		First(&user, &db.User{Name: req.Username}).
		Error; err != nil {
		return credsErr()
	}

	if ok, err := utils.VerifyPassword(req.Password, user.Password); err != nil || !ok {
		return credsErr()
	}

	cookie, err := utils.MakeCookie(CookieUser{
		ID:      user.ID,
		Name:    user.Name,
		IsAdmin: user.IsAdmin,
		Groups: utils.Map(user.Groups, func(group db.Group) string {
			return group.Name
		}),
	})
	if err != nil {
		return internalErr(err, "could not create cookie")
	}

	c.SetCookie(cookie)
	return c.String(http.StatusOK, cookie.Value)
}
