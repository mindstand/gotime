package gotime

import "time"

type LocalClock struct {
	baseTime
}

func (c LocalClock) Hour() int {
	return c.t.Hour()
}

func (c LocalClock) Minute() int {
	return c.t.Minute()
}

func (c LocalClock) Second() int {
	return c.t.Second()
}

func (c LocalClock) Nanosecond() int {
	return c.t.Nanosecond()
}

// Sub is equivalent to .GetTime().Sub()
func (c LocalClock) Sub(otherC LocalClock) time.Duration {
	return c.t.Sub(otherC.t)
}

func NewLocalClockFromTime(t time.Time) *LocalClock {
	t = t.Local()
	hour, minute, second := t.Clock()
	nano := t.Nanosecond()
	return &LocalClock{
		baseTime{
			t: time.Date(1, 1, 1, hour, minute, second, nano, time.Local),
		},
	}
}
