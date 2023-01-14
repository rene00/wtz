package tz

import (
	"path/filepath"
	"strings"

	"github.com/go-playground/tz"
)

// Tz provides info on timezone for local environment.
type Tz struct {
	// filepath to localtime which is usually /etc/localtime.
	localtimeFilepath string
}

type Zone struct {
	Name string
}

type Country struct {
	Code  string
	Name  string
	Zones []Zone
}

// ListCountries returns all Countries.
func ListCountries() []Country {
	countries := []Country{}
	for _, country := range tz.GetCountries() {
		c := Country{Name: country.Name}
		for _, zone := range country.Zones {
			c.Zones = append(c.Zones, Zone{Name: zone.Name})
		}
		countries = append(countries, c)
	}
	return countries
}

// Zoneinfo returns the local tz filename from the tz database.
func (t Tz) Zoneinfo() (string, error) {
	localtimeRelPath, err := filepath.EvalSymlinks(t.localtimeFilepath)
	if err != nil {
		return "", err
	}
	localtimeAbsPath, err := filepath.Abs(localtimeRelPath)
	if err != nil {
		return "", err
	}
	tz := localtimeAbsPath
	for _, i := range []string{"/usr/share/zoneinfo/", "/var/db/timezone/zoneinfo/", "/private/var/db/timezone/tz/2020a.1.0/zoneinfo/"} {
		tz = strings.ReplaceAll(tz, i, "")
	}
	return tz, nil
}

// NewTz accepts a localtime filepath string and returns a new tz.
func NewTz(localtimeFilepath string) *Tz {
	return &Tz{localtimeFilepath}
}
