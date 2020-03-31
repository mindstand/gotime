package gotime

import "time"

var epoch = time.Unix(0, 0).UTC()
var epochLocal = time.Date(1970, 1, 1, 0, 0, 0, 0, time.Local)

type baseTime struct {
	t time.Time
}

func (bt baseTime) GetTime() time.Time {
	return bt.t
}
