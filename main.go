package main

import (
	"flag"
	"log"

	"github.com/42LoCo42/einauth/config"
	"github.com/42LoCo42/einauth/db"
	"github.com/42LoCo42/einauth/server"
	"github.com/go-faster/errors"

	"github.com/labstack/echo/v4"
)

var (
	CONFIG config.Config
	SERVER *echo.Echo
)

func main() {
	if err := start(); err != nil {
		log.Fatal(err)
	}
}

func start() (err error) {
	address := flag.String("address", ":9001", "Address to listen on")
	configPath := flag.String("config", "einauth.yaml", "Path to config file")
	dbPath := flag.String("db", "einauth.db", "Path to database file")
	flag.Parse()

	CONFIG, err = config.Init(*configPath)
	if err != nil {
		return err
	}

	if err := db.Init(*dbPath); err != nil {
		return err
	}

	SERVER, err = server.Init()
	if err != nil {
		return err
	}

	if err := SERVER.Start(*address); err != nil {
		return errors.Wrap(err, "failed to start server")
	}

	return nil
}
