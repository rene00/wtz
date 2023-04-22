package config

import (
	"encoding/json"
	"io"
	"os"
)

type Config struct {
	// filepath is the absoluate path to the config file.
	filepath *string

	// timezones is a list of olsen timezones that will be displayed within the
	// timezone table.
	Timezones []string `json:timezones`
}

type Option func(*Config)

func WithConfigFile(f string) Option {
	return func(c *Config) {
		c.filepath = &f
	}
}

func NewConfig(options ...Option) (*Config, error) {
	c := &Config{}

	for _, option := range options {
		option(c)
	}

	if c.filepath != nil {
		f, err := os.Open(*c.filepath)
		if err != nil {
			return c, err
		}
		defer f.Close()

		if err = json.NewDecoder(io.Reader(f)).Decode(&c); err != nil {
			return c, err
		}
	}

	return c, nil
}

