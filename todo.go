package main

import (
	"flag"
	"time"

	"github.com/chooyan/todo/command"
	"github.com/chooyan/todo/util"
)

func main() {
	
	mode := flag.String("m", "list", "determin mode. list add delete are available")
	todo := flag.String("t", "", "input todo title.")
	deadline := flag.String("d", "", "input todo deadline.")
	id := flag.Int("i", -1, "input id to delete.")
	flag.Parse()

	deadlineTime, _ := time.Parse(util.SaveTimeFormat, *deadline)
	switch {
		case *mode == "list":
			command.ShowAll()
		case *mode == "add":
			command.Add(*todo, deadlineTime)
		case *mode == "delete":
			command.Delete(*id)
	}
}
