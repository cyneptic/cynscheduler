package task

import (
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
)

const (
	ImportantAndUrgent = iota + 1
	NotImportantUrgent
	ImportantNotUrgent
	NotImportantNotUrgent
)

var (
	ErrInvalidPriority = errors.New(`invalid priority, should be between 
    1 (Important and Urgent), 
    2 (Not Important and Urgent), 
    3 (Important and Not Urgent) or
    4 (Not Important and Not Urgent)
    `)

	ErrPriorityNotMatchingTime = errors.New(`incorrect priority or duration; 
    if priority is (Not Important and Not Urgent) it cannot have a duration as it is only allowed in excess time
    if priority is anything else, the duration cannot be 0 because no task is done in 0 time... Duh`)
)

type Task struct {
	Name      string        `xml:"name"`
	Desc      string        `xml:"desc"`
	ID        uuid.UUID     `xml:"id"`
	Priority  int           `xml:"priority"`
	Remaining time.Duration `xml:"remaining"`
}

func (ta *Task) GetName() string {
	return ta.Name
}

func (ta *Task) GetDesc() string {
	return ta.Desc
}

func (ta *Task) GetID() string {
	return ta.ID.String()
}

func (ta *Task) GetPrio() int {
	return ta.Priority
}

func (ta *Task) GetRemaining() time.Duration {
	return ta.Remaining
}

func NewTask(name, desc string, priority int, t time.Duration) (*Task, error) {
	if priority >= 5 || priority <= 0 {
		return nil, fmt.Errorf("received error: %w", ErrInvalidPriority)
	}

	if priority == NotImportantNotUrgent && t != 0 {
		return nil, fmt.Errorf("received error: %w", ErrPriorityNotMatchingTime)
	}

	if priority != NotImportantNotUrgent && t == 0 {
		return nil, fmt.Errorf("received error: %w", ErrPriorityNotMatchingTime)
	}

	return &Task{
		ID:        uuid.New(),
		Name:      name,
		Desc:      desc,
		Priority:  priority,
		Remaining: t,
	}, nil
}
