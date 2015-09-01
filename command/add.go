package command

import (
	"encoding/json"
	"time"

	"github.com/chooyan/todo/util"
)

// Add simply add single todo data with given title and deadline.
func Add(title string, deadline time.Time) {
	scheduleObjs := util.ReadObj()
	id := 0
	if (len(scheduleObjs.Schedules) > 0) {
		id = scheduleObjs.MaxId + 1
	} else {
		scheduleObjs = util.ScheduleList{id, []util.Schedule{}}
	}
	now := time.Now()
	schedule := util.Schedule{id, title, now.Format(util.SaveTimeFormat), deadline.Format(util.SaveTimeFormat)}
	scheduleObjs.MaxId = id
	scheduleObjs.Schedules = append(scheduleObjs.Schedules, schedule)
	scheduleJson, _ := json.Marshal(scheduleObjs)

	util.Write(string(scheduleJson))
}
