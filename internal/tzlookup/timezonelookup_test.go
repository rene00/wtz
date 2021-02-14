package tzlookup

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTzLookup(t *testing.T) {
	t.Run("checkLocaltime", func(t *testing.T) {
		var testCases = []struct {
			dirPath   string
			placeName string
			expect    string
		}{
			{
				dirPath:   "Australia",
				placeName: "Melbourne",
				expect:    "Australia/Melbourne",
			},
			{
				dirPath:   "America",
				placeName: "Thule",
				expect:    "America/Thule",
			},
			{
				dirPath:   "",
				placeName: "ROC",
				expect:    "ROC",
			},
		}
		for _, tc := range testCases {
			tempDir, err := ioutil.TempDir("", "wtz-")
			if err != nil {
				t.Fatal(err)
			}
			defer os.RemoveAll(tempDir)

			tzinfoPath := filepath.Join(tempDir, "/usr/share/zoneinfo", tc.dirPath, tc.placeName)
			tzinfoDir := filepath.Dir(tzinfoPath)

			if err := os.MkdirAll(tzinfoDir, 0755); err != nil {
				t.Fatal(err)
			}

			if err := os.Mkdir(filepath.Join(tempDir, "/etc"), 0755); err != nil {
				t.Fatal(err)
			}

			if _, err := os.Create(tzinfoPath); err != nil {
				t.Fatal(err)
			}

			if _, err := os.Create(filepath.Join(tempDir, "/usr/share/zoneinfo/UCT")); err != nil {
				t.Fatal(err)
			}

			localtimePath := filepath.Join(tempDir, "/etc/localtime")
			if err := os.Symlink(tzinfoPath, localtimePath); err != nil {
				t.Fatal(err)
			}

			tzl := New(WithLocaltimePath(localtimePath))
			got, err := tzl.checkLocaltime()
			assert.NoError(t, err)
			assert.Equal(t, tc.expect, got)
		}
	})

	t.Run("checkLocaltime with no UCT file", func(t *testing.T) {
		tempDir, err := ioutil.TempDir("", "wtz-")
		if err != nil {
			t.Fatal(err)
		}
		defer os.RemoveAll(tempDir)

		tzinfoPath := filepath.Join(tempDir, "/a/b/c/d/e/f", "test")
		tzinfoDir := filepath.Dir(tzinfoPath)

		if err := os.MkdirAll(tzinfoDir, 0755); err != nil {
			t.Fatal(err)
		}

		if err := os.Mkdir(filepath.Join(tempDir, "/etc"), 0755); err != nil {
			t.Fatal(err)
		}

		if _, err := os.Create(tzinfoPath); err != nil {
			t.Fatal(err)
		}

		localtimePath := filepath.Join(tempDir, "/etc/localtime")
		if err := os.Symlink(tzinfoPath, localtimePath); err != nil {
			t.Fatal(err)
		}

		tzl := New(WithLocaltimePath(localtimePath))
		_, err = tzl.checkLocaltime()
		assert.Error(t, err)
	})

	t.Run("checkTZ", func(t *testing.T) {
		os.Setenv("__tzlookuptest", "test1")
		defer os.Unsetenv("__tzlookuptest")
		tzl := New(WithTZName("__tzlookuptest"))
		assert.Equal(t, "test1", tzl.checkTZ())
	})
}
