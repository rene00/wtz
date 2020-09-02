package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func localAsOlson() (string, error) {
	fileLink, err := os.Readlink("/etc/localtime")
	if err != nil {
		return "", err
	}

	tz := fileLink
	for _, i := range []string{"/usr/share/zoneinfo/", "/var/db/timezone/zoneinfo/"} {
		tz = strings.ReplaceAll(tz, i, "")
	}
	return tz, nil
}

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
		for _, flag := range []string{"tz", "date"} {
			_ = viper.BindPFlag(flag, cmd.Flags().Lookup(flag))
		}
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		localOlson, err := localAsOlson()
		if err != nil {
			return err
		}

		local, err := time.LoadLocation(localOlson)
		if err != nil {
			return err
		}

		date, err := time.Parse("2006-01-02", viper.GetString("date"))
		if err != nil {
			return fmt.Errorf("Failed to correctly parse date param, want YYYY-MM-DD format\n")
		}

		tz := viper.GetStringSlice("tz")
		data, err := generateRows(date, local, tz)
		if err != nil {
			fmt.Println("DEBUG1")
			return err
		}

		_, localCity := filepath.Split(localOlson)

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
	f.StringSlice("tz", []string{""}, "timezones")
	f.String("date", time.Now().Format("2006-01-02"), "date")
}

// Execute the root command.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}
}
