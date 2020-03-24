package gotime

import "time"

type Clock struct {
	baseTime
}

func NewClock(t time.Time) *Clock {
	t = t.Local()
	hour, minute, second := t.Clock()
	nano := t.Nanosecond()
	loc := t.Location()
	return &Clock{
		baseTime{
			t: time.Date(1, 1, 1, hour, minute, second, nano, loc),
		},
	}
}
