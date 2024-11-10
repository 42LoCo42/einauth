package main

import (
	"flag"
	"log"

	"github.com/glebarez/sqlite"
	"github.com/go-faster/errors"

	// "github.com/labstack/echo/v4"
	"gorm.io/gorm"
	// "maragu.dev/gomponents"
)

type User struct {
	ID       uint
	Name     string
	Email    string
	Password string
	IsAdmin  bool
	Groups   []Group `gorm:"many2many:memberships"`
}

type Group struct {
	ID   uint
	Name string
}

func main() {
	if err := start(); err != nil {
		log.Fatal(err)
	}
}

func start() error {
	dbPath := flag.String("db", "einauth.db", "Path to database file")
	flag.Parse()

	db, err := gorm.Open(sqlite.Open(*dbPath))
	if err != nil {
		return errors.Wrap(err, "could not open DB")
	}

	if err := db.AutoMigrate(&User{}); err != nil {
		return errors.Wrap(err, "DB migration failed")
	}

	return nil
}
