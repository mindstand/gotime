package gotime

import "time"

type baseTime struct {
	t time.Time
}

func (bt baseTime) GetTime() time.Time {
	return bt.t
}
