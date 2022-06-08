package main

import (
	"github.com/marvin-hansen/goSM/scheduler_manager"
	"time"
)

func printHi() {
	println("Hi there ")
}

func printTime() {
	t := time.Now()
	println("Time is: " + t.String())
}

func main() {
	//println("Start ScheduleManager")
	//sm := scheduler_manager.NewScheduleManager()

	println("Get ScheduleManager utils")
	ut := scheduler_manager.NewSchedulerUtils()

	println("Configure min schedule")
	mCfg := ut.GetMinuteConfig(1, printHi)

	println("Get generated scheduler ID")
	mID := mCfg.SchedulerID

	println("Scheduler ID: " + mID)
}
