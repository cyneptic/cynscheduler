package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"time"

	"github.com/cyneptic/cynscheduler/app"
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

	fmt.Println(`
Press 'enter' to continue.
    `)

	fmt.Println("You pressed Enter")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	utils.Clear()
	fmt.Println("How many hours do you have remaining in this day?")
	scanner.Scan()
	n := scanner.Text()
	nInt, err := strconv.Atoi(n)
	if err != nil {
		log.Print(fmt.Errorf("%w", err))
	}

	dayRemainingTime := time.Duration(nInt) * time.Hour

	fmt.Println(`
Creating Application...
    `)

	application := app.NewApp(config, dayRemainingTime)
	go application.Start()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop
}
