package utils

import (
	"encoding/json"
	"io"
	"os"
	"path/filepath"

	"github.com/cyneptic/cynscheduler/task"
)

func LoadConfig(relPath string) ([]*task.Task, error) {
	var result []*task.Task

	configJson, err := GetConfig(relPath)
	if err != nil {
		return nil, err
	}

	parsedConfig, err := ParseConfig(configJson)
	if err != nil {
		return nil, err
	}

	for _, t := range parsedConfig {
		ta := task.NewTask(t.Name, t.Description, t.Remaining)
		result = append(result, ta)
	}

	return result, nil
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

	config, err := io.ReadAll(configFile)
	if err != nil {
		return []byte{}, err
	}

	return config, nil
}

func ParseConfig(config []byte) ([]task.TaskJSON, error) {
	var result []task.TaskJSON

	err := json.Unmarshal(config, &result)
	if err != nil {
		return []task.TaskJSON{}, err
	}

	return result, nil
}
