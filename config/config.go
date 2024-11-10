package config

import (
	"os"

	"github.com/go-faster/errors"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Domain string
	Rules  []Rule
}

type Rule struct {
	Domain   string
	Paths    []string
	Subjects []string
	Policy   string
}

func Init(path string) (config Config, err error) {
	file, err := os.Open(path)
	if err != nil {
		return config, errors.Wrap(err, "could not open config file")
	}
	defer file.Close()

	if err := yaml.NewDecoder(file).Decode(&config); err != nil {
		return config, errors.Wrap(err, "could not load config")
	}

	return config, nil
}
