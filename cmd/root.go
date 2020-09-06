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

func generateRows(date time.Time, local *time.Location, locations []string) ([][]string, error) {
	rows := [][]string{}

	for i := 0; i <= 23; i++ {
		tzRow := []string{}
		localTime, err := time.ParseInLocation("2006-01-02 15:04:05", fmt.Sprintf("%s %02d:00:00", date.Format("2006-01-02"), i), local)
		if err != nil {
			return rows, err
		}
		tzRow = append(tzRow, localTime.Format("15:04"))

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
		for _, flag := range []string{"localtime", "tz", "date", "zoneinfo", "config-file"} {
			_ = viper.BindPFlag(flag, cmd.Flags().Lookup(flag))
		}
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		c, err := config.NewConfig(cmd.Flags())
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			return err
		}

		zoneinfo := viper.GetString("zoneinfo")
		if zoneinfo == "" {
			zoneinfo = c.UserViperConfig.GetString("zoneinfo")
		}

		if zoneinfo == "" {
			t := tz.NewTz(viper.GetString("localtime"))
			zoneinfo, err = t.Zoneinfo()
			if err != nil {
				return err
			}
		}

		local, err := time.LoadLocation(zoneinfo)
		if err != nil {
			return err
		}

		date, err := time.Parse("2006-01-02", viper.GetString("date"))
		if err != nil {
			return errors.New("failed to correctly parse date param, want YYYY-MM-DD format")
		}

		tz := viper.GetStringSlice("tz")
		if len(tz) == 0 {
			tz = c.UserViperConfig.GetStringSlice("tz")
		}

		data, err := generateRows(date, local, tz)
		if err != nil {
			return err
		}

		_, localCity := filepath.Split(zoneinfo)

		cities := []string{localCity}
		for _, i := range tz {
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
	f.StringSlice("tz", []string{""}, "timezones")
	f.String("date", time.Now().Format("2006-01-02"), "date")
	f.String("zoneinfo", "", "local zoneinfo")
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
