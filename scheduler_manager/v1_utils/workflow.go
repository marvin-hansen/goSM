package v1_utils

import t "github.com/marvin-hansen/goSM/config_types"

func (u SchedulerUtils) GetMillisecondConfig(intervalUnit int, fn func()) *t.SchedulerConfig {
	return u.getConfig(t.Millisecond, intervalUnit, fn)
}

func (u SchedulerUtils) GetSecondConfig(intervalUnit int, fn func()) *t.SchedulerConfig {
	return u.getConfig(t.Second, intervalUnit, fn)
}

func (u SchedulerUtils) GetMinuteConfig(intervalUnit int, fn func()) *t.SchedulerConfig {
	return u.getConfig(t.Minute, intervalUnit, fn)
}

func (u SchedulerUtils) GetHourConfig(intervalUnit int, fn func()) *t.SchedulerConfig {
	return u.getConfig(t.Hour, intervalUnit, fn)
}

func (u SchedulerUtils) GetDayConfig(intervalUnit int, fn func()) *t.SchedulerConfig {
	return u.getConfig(t.Day, intervalUnit, fn)
}

func (u SchedulerUtils) getConfig(scheduleInterval t.ScheduleInterval, intervalUnit int, fn func()) *t.SchedulerConfig {
	return &t.SchedulerConfig{
		SchedulerID:      RandomString(14),
		ScheduleInterval: scheduleInterval,
		IntervalUnit:     intervalUnit,
		Handler:          fn,
	}
}
