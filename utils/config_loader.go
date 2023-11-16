package utils

import (
	"encoding/json"
	"io"
	"os"
	"path/filepath"

	"github.com/cyneptic/cynscheduler/task"
)

func LoadConfig(relPath string) ([]*task.Task, int, error) {
	var result []*task.Task

	configJson, err := GetConfig(relPath)
	if err != nil {
		return nil, 0, err
	}

	parsedConfig, err := ParseConfig(configJson)
	if err != nil {
		return nil, 0, err
	}

	for _, t := range parsedConfig.Tasks {
		ta := task.NewTask(t.Name, t.Description, t.Remaining, t.Important, t.Urgent)
		result = append(result, ta)
	}

	return result, parsedConfig.Hours, nil
}

func GetConfig(relPath string) ([]byte, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return []byte{}, err
	}

	configPath := filepath.Join(cwd, relPath)

	configFile, err := os.Open(configPath)
	if err != nil {
		return []byte{}, err
	}
	defer configFile.Close()

	config, err := io.ReadAll(configFile)
	if err != nil {
		return []byte{}, err
	}

	return config, nil
}

type config struct {
	Tasks []task.TaskJSON `json:"tasks"`
	Hours int             `json:"hours"`
}

func ParseConfig(cfg []byte) (config, error) {
	var result config

	err := json.Unmarshal(cfg, &result)
	if err != nil {
		return result, err
	}

	return result, nil
}
