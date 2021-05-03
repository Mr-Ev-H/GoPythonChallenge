package challenges

import (
	"fmt"
	"time"
)

// Date time looping
// Frankly, annoying that you have to google for the answer - or you're Mozarts biographer
// And who marks the CURRENT date on a calendar - surely you'd highlight important birthdays rather than the day before
// http://www.pythonchallenge.com/pc/return/uzi.html
func Solve15() {
	fromYear := 106
	toYear := 1996

	matches := 0

	for year := toYear; year >= fromYear; year -= 10 {

		checkLeap := time.Date(year, 2, 29, 0, 0, 0, 0, time.UTC)
		isLeapYear := checkLeap.Day() == 29

		if isLeapYear {

			date := time.Date(year, 1, 26, 0, 0, 0, 0, time.UTC)

			wd := date.Weekday()

			if wd == time.Monday {
				matches++
			}

			if matches == 2 {
				fmt.Printf("Now google for somebody born on: %s\n", date.Format("02 Jan 2006"))
				break
			}
		}
	}
}
