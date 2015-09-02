package main

import (
	"flag"
	"time"

	"github.com/chooyan/todo/command"
	"github.com/chooyan/todo/util"
)

func main() {
	
	mode := flag.String("m", "list", "determin mode. list add delete are available")
	before := flag.String("b", "", "show todo list before given time")
	after := flag.String("a", "", "show todo list after given time")
	todo := flag.String("t", "", "input todo title.")
	deadline := flag.String("d", "", "input todo deadline.")
	id := flag.Int("i", -1, "input id to delete.")
	flag.Parse()

	deadlineTime, _ := time.Parse(util.SaveTimeFormat, *deadline)
	switch {
		case *mode == "list":
			if len(*before) == 0 && len(*after) == 0 {
				command.ShowAll()
			} else {
				command.ShowBetween(*before, *after)
			}
			return
		case *mode == "add":
			command.Add(*todo, deadlineTime)
			return
		case *mode == "delete":
			command.Delete(*id)
			return
	}
}
