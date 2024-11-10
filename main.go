package main

import (
	"flag"
	"log"

	"github.com/42LoCo42/einauth/config"
	"github.com/42LoCo42/einauth/db"
	"github.com/42LoCo42/einauth/server"
	"github.com/go-faster/errors"
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

	if err = config.Init(*configPath); err != nil {
		return err
	}

	if err := db.Init(*dbPath); err != nil {
		return err
	}

	server, err := server.Init()
	if err != nil {
		return err
	}

	if err := server.Start(*address); err != nil {
		return errors.Wrap(err, "failed to start server")
	}

	return nil
}
