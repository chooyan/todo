package command

import (
	"fmt"
	"time"

	"github.com/chooyan/todo/util"
)

func ShowAll() {
	scheduleObjs := util.ReadObj()
	for i := 0; i < len(scheduleObjs.Schedules); i++ {
		scheduleObj := scheduleObjs.Schedules[i]
		parsedTime, _ := time.Parse(util.SaveTimeFormat, scheduleObj.Deadline)
		fmt.Println(scheduleObj.Id, "\t", parsedTime.Format(util.ShowTimeFormat), "\t", scheduleObj.Todo)
	} 
}

func ShowBetween(beforeStr string, afterStr string) {

	before, beforeErr := time.Parse(util.SaveTimeFormat, beforeStr)
	after, afterErr := time.Parse(util.SaveTimeFormat, afterStr)

	todo := util.ReadObj()
	for i := 0; i < len(todo.Schedules); i++ {
		schedule := todo.Schedules[i]
		deadline, err := time.Parse(util.SaveTimeFormat, schedule.Deadline)
		if err != nil {
			continue
		}
		if len(beforeStr) > 0 && (beforeErr != nil || deadline.After(before)) {
			continue
		}
		if len(afterStr) > 0 && (afterErr != nil || deadline.Before(after)) {
			continue
		}

		parsedTime, _ := time.Parse(util.SaveTimeFormat, schedule.Deadline)
		fmt.Println(schedule.Id, "\t", parsedTime.Format(util.ShowTimeFormat), "\t", schedule.Todo)
	}
}
