package gotime

import "time"

type LocalTime struct {
	baseTime
}

func NewLocalTime(t time.Time) *LocalTime {
	return &LocalTime{
		baseTime{
			t: t.Local(),
		},
	}
}
