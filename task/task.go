package task

import (
	"time"
)

type Timer interface {
	Tick()
	Pause()
	Toggle()
	Resume()
	Done() bool
	Timer() time.Duration
}

type Task struct {
	Timer       Timer
	Name        string
	Description string
	Spent       time.Duration
	Important   bool
	Urgent      bool
}

type TaskJSON struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Remaining   int    `json:"duration_in_minutes"`
	Important   bool   `json:"important"`
	Urgent      bool   `json:"urgent"`
}

func NewTask(name, desc string, remaining int, important, urgent bool) *Task {
	return &Task{NewTimerService(time.Duration(remaining) * time.Minute), name, desc, 0, important, urgent}
}
