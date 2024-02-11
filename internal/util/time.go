package util

import "time"

func SetLocation(t *time.Time, tz string) *time.Time {
	loc, _ := time.LoadLocation(tz)
	tLoc := t.In(loc)
	return &tLoc
}
