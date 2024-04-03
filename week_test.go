package calander

import (
	"github/com/dwethmar/calander/week"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

func TestWeeklyRecurrence_CalanderEvent(t *testing.T) {
	type fields struct {
		Days             week.Days
		ActivityStart    time.Duration
		ActivityDuration time.Duration
	}
	type args struct {
		start time.Time
		end   time.Time
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []CalanderEvent
	}{
		{
			name: "zero start should return a activity that starts at midnight",
			fields: fields{
				Days: week.Days{
					time.Monday: {},
				},
				ActivityStart:    0,
				ActivityDuration: 9 * time.Hour, // ends at 9am
			},
			args: args{
				start: time.Date(2024, 4, 1, 0, 0, 0, 0, time.UTC),
				end:   time.Date(2024, 4, 7, 0, 0, 0, 0, time.UTC),
			},
			want: []CalanderEvent{
				{
					StartTime: time.Date(2024, 4, 1, 0, 0, 0, 0, time.UTC),
					EndTime:   time.Date(2024, 4, 1, 9, 0, 0, 0, time.UTC),
				},
			},
		},
		{
			name: "activity that ends on the next day",
			fields: fields{
				Days: week.Days{
					time.Monday: {},
				},
				ActivityStart:    23 * time.Hour, // 11pm
				ActivityDuration: 2 * time.Hour,  // ends at 1am
			},
			args: args{
				start: time.Date(2024, 4, 1, 0, 0, 0, 0, time.UTC),
				end:   time.Date(2024, 4, 2, 0, 0, 0, 0, time.UTC),
			},
			want: []CalanderEvent{
				{
					StartTime: time.Date(2024, 4, 1, 23, 0, 0, 0, time.UTC),
					EndTime:   time.Date(2024, 4, 2, 1, 0, 0, 0, time.UTC),
				},
			},
		},
		{
			name: "no days",
			fields: fields{
				Days:             week.Days{},
				ActivityStart:    8 * time.Hour, // 8am
				ActivityDuration: 9 * time.Hour, // ends at 5pm
			},
			args: args{
				start: time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
				end:   time.Date(2020, 1, 2, 0, 0, 0, 0, time.UTC),
			},
			want: []CalanderEvent{},
		},
		{
			name: "one day",
			fields: fields{
				Days: map[time.Weekday]struct{}{
					time.Wednesday: {},
				},
				ActivityStart:    8 * time.Hour, // 8am
				ActivityDuration: 9 * time.Hour, // ends at 5pm
			},
			args: args{
				// The first whole week of april 2024 starts on a Monday
				start: time.Date(2024, 4, 1, 0, 0, 0, 0, time.UTC),
				end:   time.Date(2024, 4, 7, 0, 0, 0, 0, time.UTC),
			},
			want: []CalanderEvent{
				{
					StartTime: time.Date(2024, 4, 3, 8, 0, 0, 0, time.UTC),
					EndTime:   time.Date(2024, 4, 3, 17, 0, 0, 0, time.UTC),
				},
			},
		},
		{
			name: "two days",
			fields: fields{
				Days: week.Days{
					time.Wednesday: {},
					time.Friday:    {},
				},
				ActivityStart:    8 * time.Hour, // 8am
				ActivityDuration: 9 * time.Hour, // ends at 5pm
			},
			args: args{
				// The first whole week of april 2024 starts on a Monday
				start: time.Date(2024, 4, 1, 0, 0, 0, 0, time.UTC),
				end:   time.Date(2024, 4, 7, 0, 0, 0, 0, time.UTC),
			},
			want: []CalanderEvent{
				{
					StartTime: time.Date(2024, 4, 3, 8, 0, 0, 0, time.UTC),
					EndTime:   time.Date(2024, 4, 3, 17, 0, 0, 0, time.UTC),
				},
				{
					StartTime: time.Date(2024, 4, 5, 8, 0, 0, 0, time.UTC),
					EndTime:   time.Date(2024, 4, 5, 17, 0, 0, 0, time.UTC),
				},
			},
		},
		{
			name: "whole week",
			fields: fields{
				Days: week.Days{
					time.Monday:    {},
					time.Tuesday:   {},
					time.Wednesday: {},
					time.Thursday:  {},
					time.Friday:    {},
					time.Saturday:  {},
					time.Sunday:    {},
				},
				ActivityStart:    12 * time.Hour, // 12pm
				ActivityDuration: 6 * time.Hour,  // ends at 6pm
			},
			args: args{
				// The first whole week of april 2024 starts on a Monday
				start: time.Date(2024, 4, 1, 0, 0, 0, 0, time.UTC),
				end:   time.Date(2024, 4, 7, 0, 0, 0, 0, time.UTC),
			},
			want: []CalanderEvent{
				{ // Monday
					StartTime: time.Date(2024, 4, 1, 12, 0, 0, 0, time.UTC),
					EndTime:   time.Date(2024, 4, 1, 18, 0, 0, 0, time.UTC),
				},
				{ // Tuesday
					StartTime: time.Date(2024, 4, 2, 12, 0, 0, 0, time.UTC),
					EndTime:   time.Date(2024, 4, 2, 18, 0, 0, 0, time.UTC),
				},
				{ // Wednesday
					StartTime: time.Date(2024, 4, 3, 12, 0, 0, 0, time.UTC),
					EndTime:   time.Date(2024, 4, 3, 18, 0, 0, 0, time.UTC),
				},
				{ // Thursday
					StartTime: time.Date(2024, 4, 4, 12, 0, 0, 0, time.UTC),
					EndTime:   time.Date(2024, 4, 4, 18, 0, 0, 0, time.UTC),
				},
				{ // Friday
					StartTime: time.Date(2024, 4, 5, 12, 0, 0, 0, time.UTC),
					EndTime:   time.Date(2024, 4, 5, 18, 0, 0, 0, time.UTC),
				},
				{ // Saturday
					StartTime: time.Date(2024, 4, 6, 12, 0, 0, 0, time.UTC),
					EndTime:   time.Date(2024, 4, 6, 18, 0, 0, 0, time.UTC),
				},
				{ // Sunday
					StartTime: time.Date(2024, 4, 7, 12, 0, 0, 0, time.UTC),
					EndTime:   time.Date(2024, 4, 7, 18, 0, 0, 0, time.UTC),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wr := &WeeklyRecurrence{
				Weekdays:         tt.fields.Days,
				ActivityStart:    tt.fields.ActivityStart,
				ActivityDuration: tt.fields.ActivityDuration,
			}

			got := wr.CalanderEvent(tt.args.start, tt.args.end)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("CalanderEvent() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
