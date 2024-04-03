package week

import (
	"reflect"
	"testing"
	"time"
)

func TestDays_Contains(t *testing.T) {
	type args struct {
		day time.Weekday
	}
	tests := []struct {
		name string
		d    Days
		args args
		want bool
	}{
		{
			name: "contains",
			d:    New(time.Monday, time.Tuesday),
			args: args{day: time.Monday},
			want: true,
		},
		{
			name: "does not contain",
			d:    New(time.Monday, time.Tuesday),
			args: args{day: time.Wednesday},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.Contains(tt.args.day); got != tt.want {
				t.Errorf("Days.Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNew(t *testing.T) {
	type args struct {
		days []time.Weekday
	}
	tests := []struct {
		name string
		args args
		want Days
	}{
		{
			name: "empty",
			args: args{},
			want: Days{},
		},
		{
			name: "single day",
			args: args{days: []time.Weekday{time.Monday}},
			want: Days{time.Monday: {}},
		},
		{
			name: "multiple days",
			args: args{days: []time.Weekday{time.Monday, time.Tuesday}},
			want: Days{time.Monday: {}, time.Tuesday: {}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.days...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
