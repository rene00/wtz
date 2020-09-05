package tz

import (
	"path/filepath"
	"strings"
)

// Tz provides info on timezone for local environment.
type Tz struct {
	// filepath to localtime which is usually /etc/localtime.
	localtimeFilepath string
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
	for _, i := range []string{"/usr/share/zoneinfo/", "/var/db/timezone/zoneinfo/"} {
		tz = strings.ReplaceAll(tz, i, "")
	}
	return tz, nil
}

// NewTz creates a new tz.
func NewTz(localtimeFilepath string) *Tz {
	return &Tz{localtimeFilepath}
}
