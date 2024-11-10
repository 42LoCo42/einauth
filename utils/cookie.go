package utils

import (
	"crypto/rand"
	"net/http"
	"time"

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

func MakeCookie[X any](data X) (*http.Cookie, error) {
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

	return &http.Cookie{
		Name:     "einauth",
		Value:    signed,
		Expires:  end,
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
	}, nil
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
