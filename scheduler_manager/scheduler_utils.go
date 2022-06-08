package scheduler_manager

import (
	t "github.com/marvin-hansen/goSM/config_types"
	v1 "github.com/marvin-hansen/goSM/scheduler_manager/v1_utils"
)

type SchedulerUtils interface {
	GetMillisecondConfig(intervalUnit int, fn func()) *t.SchedulerConfig
	GetSecondConfig(intervalUnit int, fn func()) *t.SchedulerConfig
	GetMinuteConfig(intervalUnit int, fn func()) *t.SchedulerConfig
	GetHourConfig(intervalUnit int, fn func()) *t.SchedulerConfig
	GetDayConfig(intervalUnit int, fn func()) *t.SchedulerConfig
}

func NewSchedulerUtils() SchedulerUtils {
	return v1.NewSchedulerUtils()
}
