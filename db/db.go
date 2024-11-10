package db

import (
	"github.com/glebarez/sqlite"
	"github.com/go-faster/errors"
	"gorm.io/gorm"
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

func Init(path string) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(path))
	if err != nil {
		return nil, errors.Wrap(err, "could not open DB")
	}

	if err := db.AutoMigrate(&User{}); err != nil {
		return nil, errors.Wrap(err, "DB migration failed")
	}

	// create initial admin user if not present
	{
		var admin User
		db.
			Where(User{IsAdmin: true}).
			Attrs(User{Name: "admin"}).
			FirstOrCreate(&admin)
	}

	return db, nil
}
