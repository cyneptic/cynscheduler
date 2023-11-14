package app

import (
	"errors"
	"fmt"
	"os"
	"sort"
	"time"

	"github.com/charmbracelet/bubbles/timer"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/cyneptic/cynscheduler/task"
)

type App struct {
	CurTask    *task.Task
	Tasks      []*task.Task
	Timer      timer.Model
	IsPaused   bool
	IsFinished bool
}

func NewApp(tasks []*task.Task, remaining int) *App {
	tasks = sortListByMatrix(tasks)
	if len(tasks) == 0 {
		panic(errors.New("cannot start the program without tasks, set tasks in config.json"))
	}
	app := &App{nil, tasks, timer.New(time.Duration(remaining) * time.Hour), false, false}
	return app
}

func sortListByMatrix(tasks []*task.Task) []*task.Task {
	sort.SliceStable(tasks, func(i, j int) bool {
		if tasks[i].Important && tasks[i].Urgent && (!tasks[j].Important || !tasks[j].Urgent) {
			return true
		}

		if tasks[i].Important && !tasks[i].Urgent && !tasks[j].Important && tasks[j].Urgent {
			return true
		}

		if tasks[i].Important && !tasks[i].Urgent && !tasks[j].Important && !tasks[j].Urgent {
			return true
		}

		if !tasks[i].Important && tasks[i].Urgent && !tasks[j].Important && !tasks[j].Urgent {
			return true
		}

		return false
	})

	return tasks
}

func (a *App) Next() {
	if len(a.Tasks) == 0 {
		a.IsFinished = true
		return
	}

	a.CurTask = a.Tasks[0]
	if len(a.Tasks) > 1 {
		a.Tasks = a.Tasks[1:]
		return
	}

	a.Tasks = []*task.Task{}
}

func (a *App) Init() tea.Cmd {
	a.Next()
	return a.Timer.Init()
}

func (a *App) View() string {
	s := fmt.Sprintf("%v\n\n", a.Timer.View())
	s += fmt.Sprintf("current task: %s", a.CurTask.Name)
	if a.CurTask.Important || a.CurTask.Urgent {
		s += fmt.Sprintf(", time remaining: %v", a.CurTask.Timer.Timer())
	}
	s += "\n\n"
	if a.IsFinished {
		s = "\nGood bye ! :)\n\n"
		go func() {
			time.Sleep(1 * time.Second)
			os.Exit(0)
		}()
	}
	return s
}

type GoodByeMsg struct{}

func (a *App) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case timer.TimeoutMsg:
		a.IsFinished = true
		return a, nil

	case timer.TickMsg:
		var cmd tea.Cmd
		a.CurTask.Timer.Tick()
		a.Timer, cmd = a.Timer.Update(msg)
		if a.CurTask.Timer.Done() {
			a.Next()
		}

		return a, cmd

	case timer.StartStopMsg:
		var cmd tea.Cmd
		a.Timer, cmd = a.Timer.Update(msg)
		return a, cmd

	case tea.KeyMsg:
		switch {
		case msg.String() == "ctrl+c":
			os.Exit(1)

		case msg.Type == tea.KeySpace:
			return a, a.Timer.Toggle()
		}
	}

	return a, nil
}
