package config

import (
	"encoding/json"
	"io"
	"os"
)

type Config struct {
	// timezones is a list of olsen timezones that will be displayed within the
	// timezone table.
	Timezones []string `json:timezones`
}

type Option func(Config)

func NewConfig(filepath string, options ...Option) (Config, error) {
	c := Config{}

	f, err := os.Open(filepath)
	if err != nil {
		return c, err
	}
	defer f.Close()

	if err = json.NewDecoder(io.Reader(f)).Decode(&c); err != nil {
		return c, err
	}

	for _, option := range options {
		option(c)
	}

	return c, nil
}

