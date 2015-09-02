package util

import (
	"testing"
)

func TestRead(t *testing.T) {
	jsonstr := Read()
	if len(jsonstr) == 0 {
		t.Errorf("Read() didn't return any string")	
	}
}

func TestReadObj(t *testing.T) {
	scheduleList := ReadObj()

	if scheduleList.MaxId == 0 && len(scheduleList.Schedules) == 0 {
		t.Logf("ReadObj() returns empty object")	
	}
}
