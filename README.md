# goSM

goSM - Go Scheduler Manager. Scheduling & managing tasks. Easy, fast, & no fuzz.

## When to use

With goSM:

* you cannot add complex schedules i.e. tuesday 3 pm, except next week, and not on public holidays
* you cannot add specific time stamps i.e. every day at 12:34

If you need any of these or similar functionalities, please use gocron or any other scheduler for go.
Also, goSM uses gocron internally so if you need more advanced scheduling in goSM, please submit a PR.

However, use goSM when:

* You only need a basic regular scheduler i.e. run pullData() every 3 seconds or run analyzeLogs() every day
* You have quite a lot of those basic jobs to track, start, pause, resume and stop.
* You have to generate those jobs on the fly and thus have to ensure configuration is correct

## Install

```Bash
go get github.com/marvin-hansen/goSM
```

## Usage

Full sample code in [example file](example.go)

```Go
// Some sample functions to run 
func printHi() {
println("Hi there. I print every so often.")
}

func printTime() {
t := time.Now()
println("Time is: " + t.String())
}

func main() {
println("Start goSM")
sm := scheduler_manager.NewScheduleManager()

println("Get goSM utils")
ut := scheduler_manager.NewSchedulerUtils()

println("Configure second scheduler")
// Every three seconds, call printTime
sCfg := ut.GetSecondConfig(3, printTime)

println("Get second scheduler ID")
// Note, the ID has been generated by the util. 
// If you need a custom ID, either overwrite the ID before adding the scheduler or generate your custom config.  
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
}

func checkErr(err error) {
if err != nil {
println(err.Error())
}
}
```

## Configuration

For the bulk of standard use cases, just use the utils to generate the configuration, and you're good to go.

You can instantiate the config struct yourself and add values as you like. However, you will easily there is actually
little
wiggle room for customization other than adding custom ID strings mainly because to make configurations a simple and
robust as possible.

## Timezone support

By default, goSM uses the local Timezone. Bear in mind, your development, testing, and production system may reside
in different timezeones. In response, there is a second constructor NewScheduleManagerWithTimezone which takes a
specific timezone as an argument. If you have to run jobs at the same time in all different timezones, consider using
UTC as scheduler timezone on all those machines. Note, this decision only affects the goSM scheduler.

# Threat safety & concurrency

goSM and all its utils come with full mutex guarding of all public functions. The implementation is *very* conservative,
maybe not exactly the
best performance wise, but it takes zero risk of conflicting read/write access. In theory, this should prevent the
bulk of snafu that can happen during concurrency programming.

In practice, however, it is almost certain that you are going to run into an ordering problem long before encountering
hard concurrency problem in go.
For example, suppose you build a REST API which users can use to start and stop regular tasks. Each http handler starts
a new goroutine which then either adds, starts, pauses, resumes, stops or deletes a scheduled task. Unavoidably,
at some point a start requires will come in before the job was added and, conventionally, goSM wout raise and error
that the task does not exist.

Addressing this problem, goSM comes with a public function CheckSchedulerExists(id), which you call beforehand to verify
the scheduler is actually available before starting it. In case it's not, you can start a latch, wait a few
milliseconds,
and. check again. If the add request came in the meantime, you can start the scheduler. Otherwise, return the
appropriate error.

It is perfectly possible that a deep enough concurrency test may expose a potential problem in goSM. I cannot rule that
out, however, neither do I have the time to do a formal verification. In the mechanism of check before
start / stop with all public functions being mutex guarded has proven effective in practice.

## Why?

Many great task schedulers such as gocron exists and these tools offer fine-grained control to when and how to run a
task.
This is ideal when your task schedule can be anything from simple to mind-bending complex. However,
I had (and still have) a project in which a larger number of relatively simple jobs need to be scheduled at a regular
interval.
Nothing fancy, just a lot of minion jobs needed to be tracked. While existing libraries are wonderful at scheduling
tasks at any conceivable interval, however,
managing all those tasks could be better.

In response, I wrote a scheduler manager that tracks all those scheduled jobs, so I can add new ones, start, pause,
resume,
stop and delete any of those tasks at run-time.

While working with existing libraries, I had to debug some scheduler configurations due to incorrectly generated
scheduler config strings.
Turned out, this can happen quite easily because in many implementations the actual signature type is interface{} means
there is no way to verify the correctness of the scheduler configuration before the scheduler gets constructed. As a
result,
I've had a few instances of zombie schedulers doing nothing due to incorrect configurations that simply slipped through.

In response, I wrapped the scheduler configuration into a strongly typed struct and added a config generator to ensure
correct configs.

Ultimately, I re-wrote all my tools and utils as goSM: A Go Scheduler Manager. Scheduling & managing tasks. Easy, fast,
& no fuzz.

## Author

* Marvin Hansen
* GPG key ID: 210D39BC
* Github key ID: 369D5A0B210D39BC
* GPG Fingerprint: 4B18 F7B2 04B9 7A72 967E 663E 369D 5A0B 210D 39BC
* Public key: [key](pubkey.txt)

## Licence

* [MIT Licence](LICENSE)
* Software is "as is" without any warranty. 
