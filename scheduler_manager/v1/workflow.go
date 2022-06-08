// Copyright (c) 2021. Marvin Friedrich Lars Hansen. All Rights Reserved. Contact: marvin.hansen@gmail.com

package v1

import t "github.com/marvin-hansen/goSM/config_types"

func (c *ScheduleManager) AddScheduler(config *t.SchedulerConfig) error {
	c.state.mtx.Lock()
	err := c.addNewScheduler(config)
	c.state.mtx.Unlock()
	return err
}

func (c *ScheduleManager) CheckSchedulerExists(schedulerID string) bool {
	return c.checkSchedulerExists(schedulerID)
}

func (c *ScheduleManager) StartScheduler(schedulerID string) error {
	c.state.mtx.Lock()
	err := c.startScheduler(schedulerID)
	c.state.mtx.Unlock()
	return err

}

func (c *ScheduleManager) StopScheduler(schedulerID string) error {
	c.state.mtx.Lock()
	err := c.stopScheduler(schedulerID)
	c.state.mtx.Unlock()
	return err
}

func (c *ScheduleManager) RemoveScheduler(schedulerID string) {
	c.state.mtx.Lock()
	_ = c.stopScheduler(schedulerID)
	c.deleteScheduler(schedulerID)
	c.deleteConfig(schedulerID)
	c.state.mtx.Unlock()

}

func (c *ScheduleManager) StopAllSchedulers() {
	c.state.mtx.Lock()
	c.stopAllSchedulers()
	c.state.mtx.Unlock()
}

func (c *ScheduleManager) RemoveAllScheduler() {
	c.state.mtx.Lock()
	c.stopAllSchedulers()
	c.deleteSchedulerMap()
	c.deleteConfigMap()
	c.state.mtx.Unlock()
}
