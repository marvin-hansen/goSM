package v1_utils

import t "github.com/marvin-hansen/goSM/config_types"

func (c *SchedulerUtils) GetMillisecondConfig(intervalUnit int, fn func()) *t.SchedulerConfig {
	c.mtx.Lock()
	cfg := c.getConfig(t.Millisecond, intervalUnit, fn)
	c.mtx.Unlock()
	return cfg
}

func (c *SchedulerUtils) GetSecondConfig(intervalUnit int, fn func()) *t.SchedulerConfig {
	c.mtx.Lock()
	cfg := c.getConfig(t.Second, intervalUnit, fn)
	c.mtx.Unlock()
	return cfg
}

func (c *SchedulerUtils) GetMinuteConfig(intervalUnit int, fn func()) *t.SchedulerConfig {
	c.mtx.Lock()
	cfg := c.getConfig(t.Minute, intervalUnit, fn)
	c.mtx.Unlock()
	return cfg
}

func (c *SchedulerUtils) GetHourConfig(intervalUnit int, fn func()) *t.SchedulerConfig {
	c.mtx.Lock()
	cfg := c.getConfig(t.Hour, intervalUnit, fn)
	c.mtx.Unlock()
	return cfg
}

func (c *SchedulerUtils) GetDayConfig(intervalUnit int, fn func()) *t.SchedulerConfig {
	c.mtx.Lock()
	cfg := c.getConfig(t.Day, intervalUnit, fn)
	c.mtx.Unlock()
	return cfg
}

func (c *SchedulerUtils) getConfig(scheduleInterval t.ScheduleInterval, intervalUnit int, fn func()) *t.SchedulerConfig {
	return &t.SchedulerConfig{
		SchedulerID:      RandomString(14),
		ScheduleInterval: scheduleInterval,
		IntervalUnit:     intervalUnit,
		Handler:          fn,
	}
}
