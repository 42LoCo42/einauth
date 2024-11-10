package config

import (
	"os"

	"github.com/go-faster/errors"
	"gopkg.in/yaml.v3"
)

type Config struct {
	URL   string
	Rules []Rule
}

type Rule struct {
	Domain   string
	Paths    []string
	Subjects []string
	Policy   string
}

var CONFIG Config

func Init(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return errors.Wrap(err, "could not open config file")
	}
	defer file.Close()

	if err := yaml.NewDecoder(file).Decode(&CONFIG); err != nil {
		return errors.Wrap(err, "could not load config")
	}

	return nil
}
