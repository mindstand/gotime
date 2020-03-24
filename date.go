package gotime

import "time"

type Date struct {
	baseTime
}

func NewDate(t time.Time) *Date {
	year, month, day := t.Date()
	return &Date{
		baseTime{
			t: time.Date(year, month, day, 0, 0, 0, 0, time.UTC),
		},
	}
}
