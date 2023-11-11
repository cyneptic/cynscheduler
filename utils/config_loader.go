package utils

import (
	"encoding/xml"
	"io"
	"os"
	"strconv"
	"time"

	"github.com/cyneptic/cynscheduler/task"
)

type TaskDTO struct {
	Name      string `xml:"name"`
	Desc      string `xml:"desc"`
	Priority  string `xml:"priority"`
	Remaining string `xml:"remaining"`
}

func Config_loader() ([]*task.Task, error) {
	var config []TaskDTO

	xmlFile, err := os.Open("../config.xml")
	if err != nil {
		return nil, err
	}

	xmlData, err := io.ReadAll(xmlFile)
	if err != nil {
		return nil, err
	}

	err = xml.Unmarshal(xmlData, &config)
	if err != nil {
		return nil, err
	}

	var result []*task.Task
	for _, v := range config {
		prio, _ := strconv.Atoi(v.Priority)
		r, _ := strconv.Atoi(v.Remaining)
		remaining := time.Duration(r) * time.Minute
		task, err := task.NewTask(v.Name, v.Desc, prio, remaining)
		if err != nil {
			return nil, err
		}
		result = append(result, task)
	}

	return result, nil
}
