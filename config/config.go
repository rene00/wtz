package config

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var (
	DefaultDirName string
	viperConfig    *viper.Viper
)

type Config struct {
	UserViperConfig *viper.Viper
	FilePath        string
	Dir             string
}

// NewConfig returns a Config.
func NewConfig(flags *pflag.FlagSet) (*Config, error) {
	dir := Dir()
	configName := "wtz.json"
	configType := "json"
	if configFile, _ := flags.GetString("config-file"); configFile != "" {
		abs, err := filepath.Abs(configFile)
		if err != nil {
			return nil, err
		}
		if fmt.Sprintf("%s", filepath.Ext(abs)) != fmt.Sprintf(".%s", configType) {
			return nil, fmt.Errorf("File extension must be %s, %s", configType, configFile)
		}
		dir = filepath.Dir(abs)
		configName = filepath.Base(abs)
	}

	return &Config{
		UserViperConfig: viperConfig,
		FilePath:        filepath.Join(dir, configName),
		Dir:             dir,
	}, nil
}

// SetDefaultDirName sets DefaultDirName.
func SetDefaultDirName(binaryName string) {
	binaryNameBase := filepath.Base(binaryName)
	if binaryNameBase == "main" {
		binaryNameBase = "wtz"
	}
	DefaultDirName = strings.Replace(binaryNameBase, ".exe", "", 1)
}

// Dir returns the config dir.
func Dir() string {
	var dir string
	if runtime.GOOS == "windows" {
		dir = os.Getenv("APPDATA")
		if dir != "" {
			return filepath.Join(dir, DefaultDirName)
		}
	}

	dir = os.Getenv("WTZ_CONFIG_HOME")
	if dir != "" {
		return dir
	}

	dir = os.Getenv("XDG_CONFIG_HOME")
	if dir == "" {
		dir = filepath.Join(os.Getenv("HOME"), ".config")
	}
	if dir != "" {
		return filepath.Join(dir, DefaultDirName)
	}

	dir, _ = os.Getwd()
	return dir
}

// Save saves the config.
func (c Config) Save(basename string) error {
	if _, err := os.Stat(c.Dir); os.IsNotExist(err) {
		if err := os.MkdirAll(c.Dir, os.FileMode(0755)); err != nil {
			return err
		}
	}
	return c.UserViperConfig.WriteConfigAs(c.FilePath)
}
