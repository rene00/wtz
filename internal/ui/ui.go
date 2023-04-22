package ui

import (
	"fmt"
	"time"
)

// GenerateRows accepts a date and a list of locations and returns the list
// that will be used to generate the timezone table.
func GenerateRows(date time.Time, locations []string) ([][]string, error) {
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
