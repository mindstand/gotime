package gotime

import "time"

// Dates are always global (UTC) with no tz information. Only the day, month, and year should be accessible
type Date struct {
	baseTime
}

func (d Date) Year() int {
	return d.t.Year()
}

func (d Date) Month() time.Month {
	return d.t.Month()
}

func (d Date) Day() int {
	return d.t.Day()
}

// GetEpochDays returns the date's number of days since the unix epoch
func (d Date) GetEpochDays() int {
	return int(d.t.Sub(epoch).Hours() / 24)
}

// Sub is equivalent to .GetTime().Sub()
func (d Date) Sub(otherD Date) time.Duration {
	return d.t.Sub(otherD.t)
}

func NewDateFromTime(t time.Time) Date {
	year, month, day := t.Date()
	return Date{
		baseTime{
			t: time.Date(year, month, day, 0, 0, 0, 0, time.UTC),
		},
	}
}

func NewDate(year int, month time.Month, day int) Date {
	return Date{
		baseTime{
			t: time.Date(year, month, day, 0, 0, 0, 0, time.UTC),
		},
	}
}

// NewDateOfEpochDays is a special constructor that returns a date n days from the unix epoch
func NewDateOfEpochDays(days int) Date {
	// initialize a time at exactly the unix epoch and add n days
	t := epoch.AddDate(0, 0, days)
	return NewDateFromTime(t)
}
