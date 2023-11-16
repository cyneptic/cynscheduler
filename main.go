package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/cyneptic/cynscheduler/app"
	"github.com/cyneptic/cynscheduler/utils"
)

func main() {
	tasks, err := utils.LoadConfig("config.json")
	if err != nil {
		fmt.Println("%w\n", err)
	}

	fmt.Println("Hello, to start please enter the number of hours you have left until you go to sleep / stop working.")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n, err := strconv.Atoi(scanner.Text())
	if err != nil {
		log.Fatalf("encountered error: %v", err)
	}

	app := app.NewApp(tasks, n)

	if _, err := tea.NewProgram(app).Run(); err != nil {
		log.Fatalf("Ops, error when running program: %v", err)
	}
}
