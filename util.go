package gotime

import "time"

var epoch = time.Unix(0, 0).UTC()

type baseTime struct {
	t time.Time
}

func (bt baseTime) GetTime() time.Time {
	return bt.t
}
