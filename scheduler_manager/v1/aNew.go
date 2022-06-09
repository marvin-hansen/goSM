// Copyright (c) 2021. Marvin Friedrich Lars Hansen. All Rights Reserved. Contact: marvin.hansen@gmail.com

package v1

import "time"

type ScheduleManager struct {
	state *State
}

func NewScheduleManager() *ScheduleManager {
	comp := ScheduleManager{state: newState(nil)}
	comp.init()
	return &comp
}

func NewScheduleManagerWithTimezone(loc *time.Location) *ScheduleManager {
	comp := ScheduleManager{state: newState(loc)}
	comp.init()
	return &comp
}
