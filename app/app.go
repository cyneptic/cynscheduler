package app

import (
	"fmt"
	"sort"
	"time"

	"atomicgo.dev/keyboard"
	"atomicgo.dev/keyboard/keys"
	"github.com/cyneptic/cynscheduler/task"
	"github.com/cyneptic/cynscheduler/utils"
)

type App struct {
	Tasks         map[string]*task.Task
	CurTask       *task.Task
	Stack         []*task.Task
	RemainingTime time.Duration
	IsPaused      bool
}

func NewApp(t map[string]*task.Task, r time.Duration) *App {
	var firstTask *task.Task
	var stack []*task.Task

	for _, v := range t {
		stack = append(stack, v)
	}

	sort.Slice(stack, func(i, j int) bool {
		if stack[i].Priority == stack[j].Priority {
			return stack[i].GetRemaining() < stack[j].GetRemaining()
		}
		return stack[i].Priority > stack[j].Priority
	})

	firstTask = stack[len(stack)-1]
	stack = stack[:len(stack)-1]

	return &App{t, firstTask, stack, r, false}
}

func (a *App) Show() {
	oneStep := 100 * time.Millisecond
	for {
		if !a.IsPaused {
			fmt.Printf(`
      Current Task: %s - Remaining From Task: %s - Remaining From Day: %s
      `, a.CurTask.GetName(), a.CurTask.GetRemaining(), a.RemainingTime,
			)
			time.Sleep(oneStep)
			utils.Clear()

		} else {
			fmt.Printf(`
      Current Task: %s - Remaining From Task: %s - Remaining From Day: %s - Paused!
      `, a.CurTask.GetName(), a.CurTask.GetRemaining(), a.RemainingTime,
			)
			time.Sleep(oneStep)
			utils.Clear()
		}
	}
}

func (a *App) SyncLoop() {
	oneStep := 100 * time.Millisecond
	for a.RemainingTime > 0 {
		if !a.IsPaused {
			if a.CurTask.GetRemaining() > 0 {
				a.CurTask.Remaining -= oneStep
			}
			a.RemainingTime -= oneStep
			time.Sleep(oneStep)
		} else {
			time.Sleep(oneStep * 5)
		}
	}
}

func (a *App) ListenForPause() {
	keyboard.Listen(func(key keys.Key) (stop bool, err error) {
		if key.Code == keys.Space {
			a.IsPaused = !a.IsPaused
		}

		return false, nil // Return false to continue listening
	})
}

func (a *App) Start() {
	time.Sleep(1 * time.Second)
	a.IsPaused = false
	go a.Show()
	go a.ListenForPause()
	a.SyncLoop()
}
