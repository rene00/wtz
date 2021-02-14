// tzloookup attempts to find the local olsen timezone.

package tzlookup

import (
	"errors"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func New(opts ...Opt) (t *tzlookup) {
	t = &tzlookup{}
	t.localtimePath = "/etc/localtime"
	t.tzName = "TZ"
	for _, opt := range opts {
		opt(t)
	}
	return
}

type tzlookup struct {
	// The path of the localtime file.
	localtimePath string

	// The environment variable name where timezone could be stored.
	tzName string
}

type Opt func(*tzlookup)

func WithLocaltimePath(s string) Opt {
	return func(t *tzlookup) {
		t.localtimePath = s
	}
}

func WithTZName(s string) Opt {
	return func(t *tzlookup) {
		t.tzName = s
	}
}

func (t *tzlookup) checkTZ() (localTimezone string) {
	if v, ok := os.LookupEnv(t.tzName); ok {
		localTimezone = v
	}
	return
}

func (t *tzlookup) checkLocaltime() (string, error) {
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

func (t *tzlookup) Local() (string, error) {
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
