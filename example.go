package main

import (
	"github.com/marvin-hansen/goSM/scheduler_manager"
	"syscall"
	"time"
)

func printHi() {
	println("Hi there. I print every so often.")
}

func printTime() {
	t := time.Now()
	println("Time is: " + t.String())
}

func main() {
	println("Start ScheduleManager")
	sm := scheduler_manager.NewScheduleManager()

	println("Get ScheduleManager utils")
	ut := scheduler_manager.NewSchedulerUtils()

	println("Configure minute schedule")
	// Every one minute, call printHi
	mCfg := ut.GetMinuteConfig(1, printHi)

	println("Get generated scheduler ID")
	mID := mCfg.SchedulerID

	println("Scheduler ID: " + mID)

	println("Add minute schedule")
	err0 := sm.AddScheduler(mCfg)
	checkErr(err0)

	println("Start minute schedule")
	err1 := sm.StartScheduler(mID)
	checkErr(err1)

	println("Configure second scheduler")
	// Every three seconds, call printTime
	sCfg := ut.GetSecondConfig(3, printTime)

	println("Get second scheduler ID")
	sID := sCfg.SchedulerID

	println("Add second scheduler")
	err1 = sm.AddScheduler(sCfg)
	checkErr(err1)

	println("Start second scheduler")
	err2 := sm.StartScheduler(sID)
	checkErr(err2)

	time.Sleep(32 * time.Second)
	println("Stop second schedules")
	err3 := sm.StopScheduler(sID)
	checkErr(err3)

	println("Remove second scheduler")
	sm.RemoveScheduler(sID)

	time.Sleep(2 * time.Minute)
	println("Stop all schedulers")
	sm.StopAllSchedulers()

	println("Remove all schedulers")
	sm.RemoveAllScheduler()

	println("Done. Exit now")
	syscall.Exit(0)
}

func checkErr(err error) {
	if err != nil {
		println(err.Error())
	}
}
