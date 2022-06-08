// Copyright (c) 2021. Marvin Friedrich Lars Hansen. All Rights Reserved. Contact: marvin.hansen@gmail.com

package v1

func (c *ScheduleManager) init() {
	c.verifyInit()
}

func (c *ScheduleManager) verifyInit() {
	nilCheck(c.state, "NPE: state is nil. Fix init.")
	nilCheck(c.state.configMap, "NPE: configMap is nil. Fix init.")
	nilCheck(c.state.schedulerMap, "NPE: schedulerMap is nil. Fix init.")
}
