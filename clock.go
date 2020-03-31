package gotime

import "time"

type Clock struct {
	baseTime
}

func (c Clock) Hour() int {
	return c.t.Hour()
}

func (c Clock) Minute() int {
	return c.t.Minute()
}

func (c Clock) Second() int {
	return c.t.Second()
}

func (c Clock) Nanosecond() int {
	return c.t.Nanosecond()
}

func (c Clock) Location() *time.Location {
	return c.t.Location()
}

// Sub is equivalent to .GetTime().Sub()
func (c Clock) Sub(otherC Clock) time.Duration {
	return c.t.Sub(otherC.t)
}

func NewClockFromTime(t time.Time) Clock {
	t = t.Local()
	hour, minute, second := t.Clock()
	nano := t.Nanosecond()
	loc := t.Location()
	return Clock{
		baseTime{
			t: time.Date(1, 1, 1, hour, minute, second, nano, loc),
		},
	}
}

func NewClock(hour, minute, second, nanosecond int, location *time.Location) Clock {
	return Clock{
		baseTime{
			t: time.Date(1, 1, 1, hour, minute, second, nanosecond, location),
		},
	}
}

func NewClockOfDayNano(nanoOfDay int64) Clock {
	// todo check for time greater than one day
	return Clock{
		baseTime{
			t: epoch.Add(time.Duration(nanoOfDay)),
		},
	}
}
