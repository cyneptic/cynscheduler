package app

import (
	"errors"
	"sort"
	"time"

	"github.com/charmbracelet/bubbles/timer"
	tea "github.com/charmbracelet/bubbletea"
	styles "github.com/cyneptic/cynscheduler/style"
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
	a.CurTask = a.Tasks[1]
	if len(a.Tasks) > 1 {
		a.Tasks = a.Tasks[1:]
		return
	}

	a.Tasks = []*task.Task{}
}

func (a *App) AddRest() {
	tmp := []*task.Task{task.NewTask("Rest", "rest a bit :)", 15, true, true)}
	a.Tasks = append(tmp, a.Tasks...)
	a.CurTask = a.Tasks[0]
}

func (a *App) Init() tea.Cmd {
	a.CurTask = a.Tasks[0]
	return a.Timer.Init()
}

func (a *App) View() string {
	if a.IsFinished {
		s := "\nGood bye ! :)\n\n"
		return s
	}
	s := styles.GetTitleString(a.Timer.View())
	s += "\n\n"
	s += styles.GetStyledTable(a.Tasks)
	s += "\n\n" + styles.GetLegendString(a.CurTask, a.IsPaused)

	s = styles.GetMainStyle(s)

	return s
}

type (
	GoodByeMsg  struct{}
	FinishMsg   struct{}
	DelegateMsg struct{}
	RestMsg     struct{}
)

func (a *App) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case FinishMsg:
		if len(a.Tasks) == 1 {
			a.IsFinished = true
			return a, func() tea.Msg { return tea.QuitMsg{} }
		}
		a.Next()
	case DelegateMsg:
		if len(a.Tasks) == 1 {
			a.IsFinished = true
			return a, func() tea.Msg { return tea.QuitMsg{} }
		}
		a.Next()

	case tea.QuitMsg:
		time.Sleep(time.Second)
		return a, tea.Quit
	case timer.TimeoutMsg:
		a.IsFinished = true
		return a, func() tea.Msg { return tea.QuitMsg{} }

	case timer.TickMsg:
		var cmd tea.Cmd
		a.Timer, cmd = a.Timer.Update(msg)
		if !a.IsPaused {
			a.CurTask.Timer.Tick()
		}
		if a.CurTask.Timer.Done() {
			if len(a.Tasks) == 1 {
				a.IsFinished = true
				time.Sleep(time.Second)
				return a, func() tea.Msg { return tea.QuitMsg{} }
			}
			a.Next()
		}

		return a, cmd

	case timer.StartStopMsg:
		var cmd tea.Cmd
		a.IsPaused = !a.IsPaused
		a.Timer, cmd = a.Timer.Update(msg)
		return a, cmd

	case RestMsg:
		a.AddRest()
		return a, nil

	case tea.KeyMsg:
		switch {
		case msg.String() == "ctrl+c":
			return a, tea.Quit
		case msg.Type == tea.KeySpace:
			return a, a.Timer.Toggle()
		case msg.String() == "f":
			return a, func() tea.Msg { return FinishMsg{} }
		case msg.String() == "d":
			return a, func() tea.Msg { return DelegateMsg{} }
		case msg.String() == "r":
			return a, func() tea.Msg { return RestMsg{} }

		}
	}

	return a, nil
}
