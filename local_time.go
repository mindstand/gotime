package gotime

import "time"

type LocalTime struct {
	baseTime
}

func (lt LocalTime) Year() int {
	return lt.t.Year()
}

func (lt LocalTime) Month() time.Month {
	return lt.t.Month()
}

func (lt LocalTime) Day() int {
	return lt.t.Day()
}

func (lt LocalTime) Hour() int {
	return lt.t.Hour()
}

func (lt LocalTime) Minute() int {
	return lt.t.Minute()
}

func (lt LocalTime) Second() int {
	return lt.t.Second()
}

func (lt LocalTime) Nanosecond() int {
	return lt.t.Nanosecond()
}

// Sub is equivalent to .GetTime().Sub()
func (lt LocalTime) Sub(otherT LocalTime) time.Duration {
	return lt.t.Sub(otherT.t)
}

func NewLocalTimeFromTime(t time.Time) LocalTime {
	return LocalTime{
		baseTime{
			t: t.Local(),
		},
	}
}

func NewLocalTimeFromUnix(epochSeconds, nanoOfDay int64) LocalTime {
	t := time.Unix(epochSeconds, nanoOfDay)
	// get local offset
	_, offset := t.Zone()
	t = t.Add(time.Duration(-offset) * time.Second)

	return LocalTime{
		baseTime{
			t: t,
		},
	}
}
