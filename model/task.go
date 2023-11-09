package model

import "time"

type Task struct {
	Name              string
	Desc              string
	Applications      []string
	ID                int
	Remaining         time.Duration
	IsProcrastination bool
}
