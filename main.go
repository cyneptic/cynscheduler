package main

import (
	"fmt"
	"log"
	"time"

	"github.com/cyneptic/cynscheduler/task"
	"github.com/cyneptic/cynscheduler/utils"
)

var actionTime = 1 * time.Second

func main() {
	fmt.Println(`
Welcome to cynscheduler

Loading Config File...
    `)

	time.Sleep(actionTime)
	config, err := utils.Config_loader()
	if err != nil {
		log.Print(fmt.Errorf("%w", err))
	}

	utils.Clear()
	fmt.Println(`
Config loaded! your tasks are:
    `)

	for _, t := range config {
		fmt.Printf("%s: %s - its one of the %s tasks and you need to spend %d minutes on this task\n", t.Name, t.Desc, task.PriorityList[t.Priority], t.Remaining/time.Minute)
	}

	select {}
}
