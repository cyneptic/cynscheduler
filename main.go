package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println(`
  Welcome to Cynscheduler time management cli tool.

  the first step is to define your tasks to be tracked.
  this step will continue until you run can longer fit a certain task in your schedule
  or until you choose to add no other tasks via the 'finish' command


  to start adding tasks please enter "start"`)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	if scanner.Text() == "start" {
		fmt.Println("starting")
	} else {
		fmt.Println("???")
	}

	select {}
}
