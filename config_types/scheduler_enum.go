package config_types

type ScheduleInterval = string

const (
	Millisecond ScheduleInterval = "Millisecond"
	Second      ScheduleInterval = "Second"
	Minute      ScheduleInterval = "Minute"
	Hour        ScheduleInterval = "Hour"
	Day         ScheduleInterval = "Day"
)
