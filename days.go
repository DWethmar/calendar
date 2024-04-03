package calander

import (
	"github/com/dwethmar/calander/week"
	"time"
)

func SelectWeekdays(start, end time.Time, weekdays week.Days) []time.Time {
	var daysInPeriod []time.Time
	current := StartOfDay(start)
	for !current.After(end) {
		if weekdays.Contains(current.Weekday()) {
			daysInPeriod = append(daysInPeriod, current)
		}

		current = current.AddDate(0, 0, 1) // Move to the next day
	}

	return daysInPeriod
}
