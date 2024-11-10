package db

import (
	"github.com/42LoCo42/einauth/utils"
	"github.com/glebarez/sqlite"
	"github.com/go-faster/errors"
	"gorm.io/gorm"
)

var DB *gorm.DB

type User struct {
	ID       uint
	Name     string
	Email    string `gorm:"unique"`
	Password string
	IsAdmin  bool
	Groups   []Group `gorm:"many2many:memberships"`
}

type Group struct {
	ID   uint
	Name string `gorm:"unique"`
}

func Init(path string) (err error) {
	DB, err = gorm.Open(sqlite.Open(path))
	if err != nil {
		return errors.Wrap(err, "could not open DB")
	}

	if err := DB.AutoMigrate(&User{}); err != nil {
		return errors.Wrap(err, "DB migration failed")
	}

	// create initial admin user if not present
	{
		var admin User
		DB.
			Where(User{IsAdmin: true}).
			Attrs(User{
				Name:     "admin",
				Email:    "admin@example.org",
				Password: errors.Must(utils.HashPassword("admin")),
			}).
			FirstOrCreate(&admin)
	}

	return nil
}
