package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	now := time.Now()

	fmt.Fprint(os.Stdout, "Date\t\tEpoch / Unix timestamp\n")
	for _, row := range rows {
		fmt.Fprintf(os.Stdout, "%-10s\t%d\n", row, row.addTo(now).Unix())
	}
}

type timeUnit interface {
	fmt.Stringer
	addTo(time.Time) time.Time
}

var rows = []timeUnit{
	year(-10),
	year(-5),
	year(-3),
	year(-2),
	year(-1),
	month(-9),
	month(-6),
	month(-3),
	month(-2),
	month(-1),
	day(-28),
	day(-21),
	day(-14),
	day(-7),
	day(-3),
	day(-2),
	day(-1),
	hour(-12),
	hour(-6),
	hour(-3),
	hour(-2),
	hour(-1),
	now{},
	hour(1),
	hour(2),
	hour(3),
	hour(6),
	hour(12),
	day(1),
	day(2),
	day(3),
	day(7),
	day(14),
	day(21),
	day(28),
	month(1),
	month(2),
	month(3),
	month(6),
	month(9),
	year(1),
	year(2),
	year(3),
	year(5),
	year(10),
}

type year int

func (y year) String() string {
	if y > 0 {
		return fmt.Sprintf("In %d years", y)
	}
	return fmt.Sprintf("%d years ago", -y)
}

func (y year) addTo(t time.Time) time.Time {
	return t.AddDate(int(y), 0, 0)
}

type month int

func (m month) String() string {
	if m > 0 {
		return fmt.Sprintf("In %d months", m)
	}
	return fmt.Sprintf("%d months ago", -m)
}

func (m month) addTo(t time.Time) time.Time {
	return t.AddDate(0, int(m), 0)
}

type day int

func (d day) String() string {
	if d > 0 {
		return fmt.Sprintf("In %d days", d)
	}
	return fmt.Sprintf("%d days ago", -d)
}

func (d day) addTo(t time.Time) time.Time {
	return t.AddDate(0, 0, int(d))
}

type hour int

func (h hour) String() string {
	if h > 0 {
		return fmt.Sprintf("In %d hours", h)
	}
	return fmt.Sprintf("%d hours ago", -h)
}

func (h hour) addTo(t time.Time) time.Time {
	return t.Add(time.Duration(h) * time.Hour)
}

type now struct{}

func (now) String() string {
	return "Now"
}

func (now) addTo(t time.Time) time.Time {
	return t
}
