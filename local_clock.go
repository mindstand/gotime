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

func (c LocalClock) ToDayNano() int64 {
	return c.t.Sub(epochLocal).Nanoseconds()
}

func NewLocalClockFromTime(t time.Time) LocalClock {
	t = t.Local()
	hour, minute, second := t.Clock()
	nano := t.Nanosecond()
	return LocalClock{
		baseTime{
			t: time.Date(1970, 1, 1, hour, minute, second, nano, time.Local),
		},
	}
}

func NewLocalClock(hour, minute, second, nanosecond int) LocalClock {
	return LocalClock{
		baseTime{
			t: time.Date(1970, 1, 1, hour, minute, second, nanosecond, time.Local),
		},
	}
}

func NewLocalClockOfDayNano(nanoOfDay int64) LocalClock {
	// todo check for time greater than one day
	return LocalClock{
		baseTime{
			t: epochLocal.Add(time.Duration(nanoOfDay)),
		},
	}
}
