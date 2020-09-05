package tz

import (
	"io/ioutil"
	"os"
	"testing"
)

func tempSymlink(dest string) (string, error) {
	file, err := ioutil.TempFile("", "wtz")
	if err != nil {
		return "", err
	}
	os.Remove(file.Name())
	if err = os.Symlink(dest, file.Name()); err != nil {
		return "", err
	}
	return file.Name(), nil
}

func TestZoneinfo(t *testing.T) {
	for _, tc := range []struct {
		zoneinfo string
		want     string
	}{
		{"/usr/share/zoneinfo/Australia/Melbourne", "Australia/Melbourne"},
		{"../../../../../../../../..//usr/share/zoneinfo/Australia/ACT", "Australia/ACT"},
		{"/usr/share/zoneinfo/NZ", "NZ"},
	} {
		s, err := tempSymlink(tc.zoneinfo)
		if err != nil {
			t.Error(err)
		}
		defer os.Remove(s)

		_t := &Tz{s}
		got, err := _t.Zoneinfo()
		if err != nil {
			t.Error(err)
		}
		if got != tc.want {
			t.Errorf("want %s but got %s", tc.want, got)
		}
	}
}
