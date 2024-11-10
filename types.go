package main

///// stored /////

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

///// rules /////

type Rule struct {
	Domain   string   `yaml:"domain"`
	Paths    []string `yaml:"paths"`
	Subjects []string `yaml:"subjects"`
}
