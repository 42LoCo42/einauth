package utils

import (
	"crypto/rand"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/42LoCo42/einauth/config"
	"github.com/go-faster/errors"
	"github.com/golang-jwt/jwt/v5"
)

var (
	JWT_METHOD = jwt.SigningMethodHS512
	JWT_SECRET = make([]byte, 64)
)

func init() {
	errors.Must(rand.Read(JWT_SECRET))
}

type CookieData[X any] struct {
	jwt.RegisteredClaims
	Data X
}

func MakeCookie(name, value string) *http.Cookie {
	return &http.Cookie{
		Name:     fmt.Sprintf("einauth-%s", name),
		Value:    value,
		Domain:   config.CONFIG.Domain,
		Path:     "/",
		Secure:   strings.HasPrefix(config.CONFIG.URL, "https"),
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
	}
}

func ClearCookie(name string) *http.Cookie {
	cookie := MakeCookie(name, "0")
	cookie.Expires = time.Unix(0, 0)
	return cookie
}

func SignCookie[X any](data X) (*http.Cookie, error) {
	now := time.Now()
	end := now.Add(time.Hour * 24)

	token := jwt.NewWithClaims(
		JWT_METHOD,
		CookieData[X]{
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(end),
				IssuedAt:  jwt.NewNumericDate(now),
				Issuer:    "einauth",
			},
			Data: data,
		},
	)

	signed, err := token.SignedString(JWT_SECRET)
	if err != nil {
		return nil, errors.Wrap(err, "could not sign cookie")
	}

	cookie := MakeCookie("token", signed)
	cookie.Expires = end
	return cookie, nil
}

func VerifyCookie[X any](cookie *http.Cookie) (*X, error) {
	var data CookieData[X]
	_, err := jwt.ParseWithClaims(
		cookie.Value,
		&data,
		func(t *jwt.Token) (any, error) {
			if t.Method != JWT_METHOD {
				return nil, errors.Errorf(
					"invalid signing method %s",
					t.Method.Alg(),
				)
			}

			return JWT_SECRET, nil
		},
	)
	if err != nil {
		return nil, errors.Wrap(err, "could not verify cookie")
	}

	return &data.Data, nil
}
