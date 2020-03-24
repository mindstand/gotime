package gotime

import "time"

type LocalClock struct {
	baseTime
}

func NewLocalClock(t time.Time) *LocalClock {
	t = t.Local()
	hour, minute, second := t.Clock()
	nano := t.Nanosecond()
	return &LocalClock{
		baseTime{
			t: time.Date(1, 1, 1, hour, minute, second, nano, time.Local),
		},
	}
}
