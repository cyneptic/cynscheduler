package task

import "time"

type Task struct {
	Name        string
	Description string
	Remaining   time.Duration
	Spent       time.Duration
}

type TaskJSON struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Remaining   int    `json:"duration_of_task"`
}

func NewTask(name, desc string, remaining int) *Task {
	return &Task{name, desc, time.Duration(remaining) * time.Minute, 0}
}
