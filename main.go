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

	// create initial admin user if not present
	{
		var admin User
		if err := db.First(&admin, &User{IsAdmin: true}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
			// admin:admin
			admin = User{
				Name:     "admin",
				Password: "$argon2id$v=19$m=65536,t=3,p=4$0YzwtVN53P2Dm8hw4pSjmA$KQazfll8MamwA+D1gDI9pEZ2TLM/6tjn4RGMib/rw3M",
				IsAdmin:  true,
			}

			if err := db.Create(&admin).Error; err != nil {
				return errors.Wrap(err, "could not create initial admin user")
			}

			log.Print("Initial admin user created with creds `admin:admin`")
		}
	}

	return nil
}
