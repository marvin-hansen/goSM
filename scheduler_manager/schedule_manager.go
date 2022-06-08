// Copyright (c) 2021. Marvin Friedrich Lars Hansen. All Rights Reserved. Contact: marvin.hansen@gmail.com

package scheduler_manager

import (
	t "github.com/marvin-hansen/goSM/config_types"
	v1 "github.com/marvin-hansen/goSM/scheduler_manager/v1"
)

type ScheduleManager interface {
	AddScheduler(config *t.SchedulerConfig) error
	CheckSchedulerExists(schedulerID string) bool
	StartScheduler(schedulerID string) error
	StopScheduler(schedulerID string) error
	RemoveScheduler(schedulerID string)
	StopAllSchedulers()
	RemoveAllScheduler()
}

// NewScheduleManager creates a new schedule manager and starts an initial scheduler.
// Usage:
//  sm := scheduler_manager.NewScheduleManager()
//  ut := scheduler_manager.NewSchedulerUtils()
// 	cfg := ut.GetMinuteConfig(1, printHi)
//  id := mCfg.SchedulerID
//
// 	sm.AddScheduler(cfg)
//  sm.StartScheduler(id)
//  sm.StopScheduler(id)
// 	sm.RemoveScheduler(id)
func NewScheduleManager() ScheduleManager {
	return v1.NewScheduleManager()
}
