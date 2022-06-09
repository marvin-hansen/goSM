package v1_utils

import "sync"

type SchedulerUtils struct {
	mtx sync.RWMutex
}

func NewSchedulerUtils() SchedulerUtils {
	return SchedulerUtils{}
}
