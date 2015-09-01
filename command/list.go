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

func Show(id string) {
}
