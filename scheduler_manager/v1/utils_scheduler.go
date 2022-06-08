package v1

import "github.com/go-co-op/gocron"

func (c *ScheduleManager) addScheduler(schedulerID string, scheduler *gocron.Scheduler) {
	c.state.schedulerMap[schedulerID] = scheduler
}

func (c *ScheduleManager) getScheduler(schedulerID string) *gocron.Scheduler {
	exists := c.checkSchedulerExists(schedulerID)
	if !exists {
		return nil
	}
	return c.state.schedulerMap[schedulerID]
}

func (c *ScheduleManager) deleteScheduler(schedulerID string) {
	exists := c.checkSchedulerExists(schedulerID)
	if !exists { //doesn't exists, do nothing
		return
	}

	delete(c.state.schedulerMap, schedulerID)
}

func (c *ScheduleManager) deleteSchedulerMap() {
	c.state.schedulerMap = make(map[string]*gocron.Scheduler)
}

func (c *ScheduleManager) checkSchedulerExists(schedulerID string) (exists bool) {
	if _, ok := c.state.schedulerMap[schedulerID]; ok {
		return true
	}
	return false
}

func (c *ScheduleManager) isSchedulerMapEmpty() bool {
	if len(c.state.schedulerMap) == 0 {
		return true
	}
	return false
}
