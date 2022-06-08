// Copyright (c) 2021. Marvin Friedrich Lars Hansen. All Rights Reserved. Contact: marvin.hansen@gmail.com

package v1

type ScheduleManager struct {
	state *State
}

func NewScheduleManager() *ScheduleManager {
	comp := ScheduleManager{state: newState()}
	comp.init()
	return &comp
}
