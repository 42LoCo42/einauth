package main

import (
	"github.com/go-faster/errors"
	"github.com/matthewhartstonge/argon2"
)

func HashPassword(pass string) (string, error) {
	argon := argon2.DefaultConfig()

	hash, err := argon.HashEncoded([]byte(pass))
	if err != nil {
		return "", errors.Wrap(err, "could not hash password")
	}

	return string(hash), nil
}

func VerifyPassword(pass, hash string) (bool, error) {
	ok, err := argon2.VerifyEncoded([]byte(pass), []byte(hash))
	if err != nil {
		return false, errors.Wrap(err, "could not verify password")
	}

	return ok, nil
}
