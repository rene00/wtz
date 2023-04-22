// tzloookup attempts to find the local olsen timezone.

package tzlookup

import (
	"errors"
	"os"
	"path"
	"path/filepath"
	"strings"
)

// New creates a new Tzlookup.
func New(localtime string, opts ...Opt) (t *Tzlookup) {
	t = &Tzlookup{localtimePath: localtime, tzName: "TZ"}
	for _, opt := range opts {
		opt(t)
	}
	return
}

// Tzlookup provides methods to lookup the local timezone.
type Tzlookup struct {
	// The path of the localtime file.
	localtimePath string

	// The environment variable name where timezone could be stored.
	tzName string
}

// Opt is a functional option for Tzlookup.
type Opt func(*Tzlookup)

// WithTZName is a function option that provides a way for the called to set the environment variable name used to lookup the local timezone.
func WithTZName(s string) Opt {
	return func(t *Tzlookup) {
		t.tzName = s
	}
}

func (t *Tzlookup) checkTZ() (localTimezone string) {
	if v, ok := os.LookupEnv(t.tzName); ok {
		localTimezone = v
	}
	return
}

func (t *Tzlookup) checkLocaltime() (string, error) {
	absPath, err := filepath.EvalSymlinks(t.localtimePath)
	if err != nil {
		return "", err
	}
	localTimezone := filepath.Base(absPath)
	localTimezoneDir := filepath.Dir(absPath)
	tzslice := []string{localTimezone}
	dirCount := strings.Count(localTimezoneDir, "/")

	for i := 1; i < dirCount; i++ {
		uctFile := path.Join(localTimezoneDir, "UCT")
		if i == (dirCount - 1) {
			return "", errors.New("Unable to find local timezone")
		}
		if _, err := os.Stat(uctFile); os.IsNotExist(err) {
			tzslice = append(tzslice, filepath.Base(localTimezoneDir))
			localTimezoneDir = filepath.Dir(localTimezoneDir)
			continue
		}
		break
	}

	for i, j := 0, len(tzslice)-1; i < j; i, j = i+1, j-1 {
		tzslice[i], tzslice[j] = tzslice[j], tzslice[i]
	}

	localTimezone = strings.Join(tzslice, "/")
	return localTimezone, nil
}

// Local attempts to lookup and return the local timezone.
func (t *Tzlookup) Local() (string, error) {
	var err error
	localTimezone := t.checkTZ()
	if localTimezone == "" {
		localTimezone, err = t.checkLocaltime()
		if err != nil {
			return "", err
		}
	}
	return localTimezone, nil
}
