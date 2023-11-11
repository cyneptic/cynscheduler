package app

import (
	"fmt"
	"time"

	"github.com/cyneptic/cynscheduler/task"
)

type App struct {
	Tasks         map[string]*task.Task
	RemainingTime time.Duration
	IsPaused      bool
}

func NewApp(t map[string]*task.Task, r time.Duration) *App {
	return &App{t, r, false}
}

func (a *App) Start() {
	a.IsPaused = false
	for a.RemainingTime > 0 {
		if !a.IsPaused {
			oneStep := 100 * time.Millisecond
			a.RemainingTime -= oneStep
			time.Sleep(oneStep)
		}
	}
}

func (a *App) Pause() {
	a.IsPaused = true
}

func (a *App) Resume() {
	a.IsPaused = false
}

func (a *App) SwitchMode() {
	a.Pause()

	i := 1
	for name := range a.Tasks {
		fmt.Printf("%d. %s", i, name)
		i++
	}
	fmt.Print("select your current task: ")
}
