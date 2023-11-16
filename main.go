package main

import (
	"fmt"
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/cyneptic/cynscheduler/app"
	"github.com/cyneptic/cynscheduler/utils"
)

func main() {
	tasks, n, err := utils.LoadConfig("config.json")
	if err != nil {
		fmt.Println("%w\n", err)
	}

	app := app.NewApp(tasks, n)

	if _, err := tea.NewProgram(app).Run(); err != nil {
		log.Fatalf("Ops, error when running program: %v", err)
	}
}
