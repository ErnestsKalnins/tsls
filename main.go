package main

import (
	"fmt"
	"os"
	"time"
)

// period contains the time increments used in the table. It is used as a sum
// type.
type period struct {
	years, months, days, hours int
}

// String returns a human readable string description of the period.
func (p period) String() string {
	if p.hours != 0 {
		if p.hours > 0 {
			return fmt.Sprintf("In %d hours", p.hours)
		} else {
			return fmt.Sprintf("%d hours ago", -p.hours)
		}
	}

	if p.days != 0 {
		if p.days > 0 {
			return fmt.Sprintf("In %d days", p.days)
		} else {
			return fmt.Sprintf("%d days ago", -p.days)
		}
	}

	if p.months != 0 {
		if p.months > 0 {
			return fmt.Sprintf("In %d months", p.months)
		} else {
			return fmt.Sprintf("%d months ago", -p.months)
		}
	}

	if p.years != 0 {
		if p.years > 0 {
			return fmt.Sprintf("In %d years", p.years)
		} else {
			return fmt.Sprintf("%d years ago", -p.years)
		}
	}

	return "Now"
}

func (p period) addTo(t time.Time) time.Time {
	if p.hours != 0 {
		return t.Add(time.Duration(p.hours) * time.Hour)
	} else {
		return t.AddDate(p.years, p.months, p.days)
	}
}

func main() {
	now := time.Now()

	fmt.Fprint(os.Stdout, "Date\t\tEpoch / Unix timestamp\n")
	for _, row := range rows {
		fmt.Fprintf(os.Stdout, "%-10s\t%d\n", row, row.addTo(now).Unix())
	}
}

var rows = []period{
	{years: -10},
	{years: -5},
	{years: -3},
	{years: -2},
	{years: -1},
	{months: -9},
	{months: -6},
	{months: -3},
	{months: -2},
	{months: -1},
	{days: -28},
	{days: -21},
	{days: -14},
	{days: -7},
	{days: -3},
	{days: -2},
	{days: -1},
	{hours: -12},
	{hours: -6},
	{hours: -3},
	{hours: -2},
	{hours: -1},
	{},
	{hours: 1},
	{hours: 2},
	{hours: 3},
	{hours: 6},
	{hours: 12},
	{days: 1},
	{days: 2},
	{days: 3},
	{days: 7},
	{days: 14},
	{days: 21},
	{days: 28},
	{months: 1},
	{months: 2},
	{months: 3},
	{months: 6},
	{months: 9},
	{years: 1},
	{years: 2},
	{years: 3},
	{years: 5},
	{years: 10},
}
