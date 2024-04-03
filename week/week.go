package week

import "time"

// Days is a set of days of the week
type Days map[time.Weekday]struct{}

// Contains returns true if the day is in the set
func (d Days) Contains(day time.Weekday) bool {
	_, ok := d[day]
	return ok
}

func New(days ...time.Weekday) Days {
	d := make(Days)
	for _, day := range days {
		d[day] = struct{}{}
	}
	return d
}
