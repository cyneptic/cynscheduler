package utils

import (
	"encoding/json"
	"io"
	"os"
	"time"

	"github.com/cyneptic/cynscheduler/task"
)

type TaskDTO struct {
	Desc      string `json:"desc"`
	Priority  int    `json:"priority"`
	Remaining int    `json:"remaining"`
}

func Config_loader() ([]*task.Task, error) {
	config := make(map[string]TaskDTO)

	configFile, err := os.Open("config.json")
	if err != nil {
		return nil, err
	}

	configData, err := io.ReadAll(configFile)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(configData, &config)
	if err != nil {
		return nil, err
	}

	var result []*task.Task

	for name, v := range config {
		task, err := task.NewTask(name, v.Desc, v.Priority, time.Duration(v.Remaining)*time.Minute)
		if err != nil {
			return nil, err
		}
		result = append(result, task)
	}

	return result, nil
}
