package cmd

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/olekukonko/tablewriter"
	"github.com/rene00/wtz/config"
	"github.com/rene00/wtz/internal/tz"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func generateRows(date time.Time, locations []string) ([][]string, error) {
	rows := [][]string{}

	primaryTimezone, err := time.LoadLocation(locations[0])
	if err != nil {
		return rows, err
	}

	for i := 0; i <= 23; i++ {
		tzRow := []string{}
		localTime, err := time.ParseInLocation("2006-01-02 15:04:05", fmt.Sprintf("%s %02d:00:00", date.Format("2006-01-02"), i), primaryTimezone)
		if err != nil {
			return rows, err
		}

		for _, i := range locations {
			tzLoc, err := time.LoadLocation(i)
			if err != nil {
				return rows, err
			}
			tzRow = append(tzRow, localTime.In(tzLoc).Format("15:04"))
		}

		rows = append(rows, tzRow)
	}

	return rows, nil
}

var rootCmd = &cobra.Command{
	Use: "wtz [command]",
	PreRun: func(cmd *cobra.Command, args []string) {
		for _, flag := range []string{"localtime", "timezones", "date", "include-local-timezone", "config-file"} {
			_ = viper.BindPFlag(flag, cmd.Flags().Lookup(flag))
		}
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		c, err := config.NewConfig(cmd.Flags())
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			return err
		}

		var localTimezone string

		if viper.GetBool("include-local-timezone") {
			t := tz.NewTz(viper.GetString("localtime"))
			localTimezone, err = t.Zoneinfo()
			if err != nil {
				return err
			}
		}

		date, err := time.Parse("2006-01-02", viper.GetString("date"))
		if err != nil {
			return errors.New("failed to correctly parse date param, want YYYY-MM-DD format")
		}

		timezones := viper.GetStringSlice("timezones")
		if len(timezones) == 0 {
			timezones = c.UserViperConfig.GetStringSlice("timezones")
		}

		if len(timezones) == 0 {
			return errors.New("No timezones set")
		}

		if viper.GetBool("include-local-timezone") {
			timezones = append([]string{localTimezone}, timezones...)
		}

		data, err := generateRows(date, timezones)
		if err != nil {
			return err
		}

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
		return nil
	},
}

func init() {
	f := rootCmd.Flags()

	f.String("localtime", "/etc/localtime", "filepath to localtime which is usually /etc/localtime")
	f.StringSlice("timezones", []string{""}, "A comma separated list of timezones")
	f.String("date", time.Now().Format("2006-01-02"), "date")
	f.Bool("include-local-timezone", true, "Include local timezone")
	f.String("config-file", "", "config file")
}

// Execute the root command.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	BinaryName := os.Args[0]
	config.SetDefaultDirName(BinaryName)
}
