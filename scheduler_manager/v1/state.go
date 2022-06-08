// Copyright (c) 2021. Marvin Friedrich Lars Hansen. All Rights Reserved. Contact: marvin.hansen@gmail.com

package v1

import (
	"github.com/go-co-op/gocron"
	t "github.com/marvin-hansen/goSM/config_types"
	"sync"
)

type State struct {
	mtx          sync.RWMutex
	configMap    map[string]*t.SchedulerConfig
	schedulerMap map[string]*gocron.Scheduler
}

func newState() (state *State) {
	state = &State{
		configMap:    make(map[string]*t.SchedulerConfig),
		schedulerMap: make(map[string]*gocron.Scheduler),
	}
	return state
}
