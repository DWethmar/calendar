package calander

import "time"

type CalanderEvent struct {
	StartTime time.Time
	EndTime   time.Time
}

type Recurrence interface {
	CalanderEvents(activity CalanderEvent, start, end time.Time) []CalanderEvent
}
