package command

import(
	"encoding/json"
	
	"github.com/chooyan/todo/util"
)

func Delete(id int) {
	scheduleList := util.ReadObj()

	newScheduleList := util.ScheduleList{scheduleList.MaxId, []util.Schedule{}}
	for i := 0; i < len(scheduleList.Schedules); i++ {
		if id != scheduleList.Schedules[i].Id {
			newScheduleList.Schedules = append(newScheduleList.Schedules, scheduleList.Schedules[i])	
		}
	}

	scheduleJson, _ := json.Marshal(newScheduleList)
	util.Write(string(scheduleJson))	
}
