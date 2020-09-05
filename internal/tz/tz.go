package tz

import (
	"path/filepath"
	"strings"
)

type tz struct {
	// filepath to localtime which is usually /etc/localtime.
	localtimeFilepath string
}

// Zoneinfo returns the local tz filename from the tz database.
func (t tz) Zoneinfo() (string, error) {
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

func NewTz(localtimeFilepath string) *tz {
	return &tz{localtimeFilepath}
}
