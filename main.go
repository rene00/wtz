package main

import (
	"flag"
	"fmt"
	"os"
	"os/user"
	"path"
	"path/filepath"
	"strings"
	"time"

	"github.com/olekukonko/tablewriter"
	"github.com/rene00/wtz/config"
	"github.com/rene00/wtz/internal/tzlookup"
	"github.com/rene00/wtz/internal/ui"
)

const version = "0.1.2"

/*

$ ./wtz --help
Usage:
  wtz [command] [flags]
  wtz [command]

Available Commands:
  help        Help about any command
  version     show version

Flags:
      --config-file string       config file
      --date string              date (default "2023-04-22")
  -h, --help                     help for wtz
      --include-local-timezone   Include local timezone (default true)
      --localtime string         filepath to localtime which is usually /etc/localtime (default "/etc/localtime")
      --timezones strings        A comma separated list of timezones
*/

type stringList []string

func (s *stringList) String() string {
	return fmt.Sprintf("%v", *s)
}

func (s *stringList) Set(value string) error {
	*s = strings.Split(value, ",")
	return nil
}

type flags struct {
	date                 *string
	configFile           *string
	timezones            *stringList
	includeLocalTimezone *bool
}

func main() {
	cmd := flag.NewFlagSet("wtz", flag.ExitOnError)

	flags := flags{}
	flags.date = cmd.String("date", time.Now().Format("2006-01-02"), fmt.Sprintf("The date to display (default: %s)", time.Now().Format("2006-01-02")))
	flags.includeLocalTimezone = cmd.Bool("include-local-timezone", true, "Include the local timezone (default: true)")

	usr, err := user.Current()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
	flags.configFile = cmd.String("config-file", path.Join(usr.HomeDir, ".config", "wtz", "wtz.json"), "Config file (defaults to ~/.config/wtz/wtz.json)")

	timezonesList := stringList{}
	flags.timezones = &timezonesList
	cmd.Var(&timezonesList, "timezones", "A comma separated list of timezones")

	cmd.Parse(os.Args[1:])

	c, err := config.NewConfig(*flags.configFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	date, err := time.Parse("2006-01-02", *flags.date)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to correctly parse date param, want YYYY-MM-DD format\n")
		os.Exit(1)
	}

	timezones := stringList{}

	if len(c.Timezones) > 0 {
		for _, timezone := range c.Timezones {
			timezones = append(timezones, timezone)
		}
	}

	if len(*flags.timezones) > 0 {
		timezones = *flags.timezones
	}

	if *flags.includeLocalTimezone {
		var localTimezone string
		tzl := tzlookup.New()
		localTimezone, err = tzl.Local()
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
		}
		timezones = append([]string{localTimezone}, timezones...)
	}

	data, err := ui.GenerateRows(date, timezones)
	cities := []string{}
	for _, i := range timezones {
		_, city := filepath.Split(i)
		cities = append(cities, city)
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(cities)
	for _, v := range data {
		table.Append(v)
	}
	table.Render()
}
