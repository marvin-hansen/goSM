package config_types

type SchedulerConfig struct {
	// Example: Call every one minute function PrintStatus()
	SchedulerID      string           // i.e. PrintStatusOncePerMinute
	ScheduleInterval ScheduleInterval // Minute
	IntervalUnit     int              // 1
	Handler          func()           // PrintStatus()
}
