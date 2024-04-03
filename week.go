package calander

import (
	"github/com/dwethmar/calander/week"
	"time"
)

type WeeklyRecurrence struct {
	Weekdays         week.Days
	ActivityStart    time.Duration // Time after the start of the day when the activity begins
	ActivityDuration time.Duration // How long the activity lasts
}

func (wr *WeeklyRecurrence) CalanderEvent(start, end time.Time) []CalanderEvent {
	if wr.ActivityStart > 24*time.Hour {
		return nil
	}

	daysInRange := SelectWeekdays(start, end, wr.Weekdays)
	var activities = make([]CalanderEvent, len(daysInRange))

	for i := range daysInRange {
		activities[i] = CalanderEvent{
			StartTime: daysInRange[i].Add(wr.ActivityStart),
			EndTime:   daysInRange[i].Add(wr.ActivityStart + wr.ActivityDuration),
		}
	}

	return activities
}
