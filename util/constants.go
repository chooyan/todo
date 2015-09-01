package util

import (
	"path"
        "github.com/kardianos/osext"	
)

const (
	SaveTimeFormat = "200601021504"
	ShowTimeFormat = "2006/01/02 15:04"
	SaveFileName   = "todo.json"
)

type ScheduleList struct {
	MaxId int
	Schedules []Schedule	
}

type Schedule struct {
        Id int 
        Todo string
        RegisterTime string
        Deadline string
}

func GetFilePath() string {
	return path.Join(path.Dir(GetFolderPath()), "bin", "res", "todo", SaveFileName)
}

func GetFolderPath() string {
        folderPath, _ := osext.ExecutableFolder()
	return folderPath
}
