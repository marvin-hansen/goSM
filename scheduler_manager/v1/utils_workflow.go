package v1

import (
	"errors"
	"github.com/go-co-op/gocron"
	t "github.com/marvin-hansen/goSM/config_types"
)

func (c *ScheduleManager) addNewScheduler(config *t.SchedulerConfig) error {
	exists := c.checkConfigExists(config.SchedulerID)
	if exists {
		return errors.New("Scheduler with the ID already exists" + config.SchedulerID)
	}

	s := c.getConfiguredScheduler(*config)
	c.addScheduler(config.SchedulerID, s)
	c.addConfig(config)
	return nil
}

func (c *ScheduleManager) startScheduler(schedulerID string) error {
	scheduler := c.getScheduler(schedulerID)
	if scheduler == nil {
		return errors.New("No scheduler found for id: " + schedulerID)
	}

	scheduler.StartAsync()
	return nil
}

func (c *ScheduleManager) stopScheduler(schedulerID string) error {
	s := c.getScheduler(schedulerID)
	if s == nil {
		return errors.New("No scheduler found for schedulerID: " + schedulerID)
	}
	s.Stop()
	c.deleteScheduler(schedulerID)
	return nil
}

func (c *ScheduleManager) stopAllSchedulers() {
	for _, s := range c.state.schedulerMap {
		if s != nil {
			s.Stop()
		}
	}
}

func (c *ScheduleManager) getConfiguredScheduler(config t.SchedulerConfig) *gocron.Scheduler {
	schedule := config.ScheduleInterval
	interval := config.IntervalUnit
	fn := config.Handler

	switch schedule {
	case t.Millisecond:
		s := c.getNewScheduler()
		_, _ = s.Every(interval).Millisecond().Do(fn)
		return s

	case t.Second:
		s := c.getNewScheduler()
		_, _ = s.Every(interval).Second().Do(fn)
		return s

	case t.Minute:
		s := c.getNewScheduler()
		_, _ = s.Every(interval).Minute().Do(fn)
		return s

	case t.Hour:
		s := c.getNewScheduler()
		_, _ = s.Every(interval).Hour().Do(fn)
		return s

	case t.Day:
		s := c.getNewScheduler()
		_, _ = s.Every(interval).Day().Do(fn)
		return s

	default:
		return nil
	}
}

func (c *ScheduleManager) getNewScheduler() *gocron.Scheduler {
	return gocron.NewScheduler(c.state.loc)
}
