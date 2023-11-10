package utils

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/cyneptic/cynscheduler/task"
)

var (
	ErrEmptyName             = errors.New("name cannot be empty")
	ErrEmptyDesc             = errors.New("description cannot be empty")
	ErrNegativeDurationInput = errors.New("duration of a task cannot be negative minutes.. hello?")
)

func ValidateNameAndDescription(name, desc string) error {
	if name == "" {
		return fmt.Errorf("received error: %w", ErrEmptyName)
	}

	if desc == "" {
		return fmt.Errorf("received error: %w", ErrEmptyDesc)
	}

	return nil
}

func ValidatePriority(priority string) (int, error) {
	n, err := strconv.Atoi(priority)
	if err != nil {
		return -1, fmt.Errorf("received error: %w", err)
	}

	if !In(n, []int{task.ImportantAndUrgent, task.NotImportantUrgent, task.NotImportantNotUrgent, task.ImportantNotUrgent}) {
		return -1, fmt.Errorf("received error: %w", task.ErrInvalidPriority)
	}

	return n, nil
}

func ValidateDuration(s string) (time.Duration, error) {
	n, err := strconv.Atoi(s)
	if err != nil {
		return -1, fmt.Errorf("received error: %w", err)
	}

	if n < 0 {
		return -1, fmt.Errorf("received error: %w", ErrNegativeDurationInput)
	}

	t := time.Duration(n) * time.Minute
	return t, nil
}

func In(n int, i []int) bool {
	for _, v := range i {
		if n == v {
			return true
		}
	}
	return false
}
