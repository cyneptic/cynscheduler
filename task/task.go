package task

import (
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
)

const (
	ImportantAndUrgent = iota + 1
	ImportantNotUrgent
	NotImportantUrgent
	NotImportantNotUrgent
)

var (
	errInvalidPriority = errors.New(`invalid priority, should be between 
    1 (Important and Urgent), 
    2 (Important and Not Urgent),
    3 (Not Important and Urgent) or
    4 (Not Important and Not Urgent)
    `)

	errPriorityNotMatchingTime = errors.New(`incorrect priority or duration; 
    if priority is (Not Important and Not Urgent) it cannot have a duration as it is only allowed in excess time
    if priority is anything else, the duration cannot be 0 because no task is done in 0 time... Duh`)
)

type Task struct {
	name      string
	desc      string
	apps      []string
	id        uuid.UUID
	priority  int
	remaining time.Duration
}

func (ta *Task) GetName() string {
	return ta.name
}

func (ta *Task) GetDesc() string {
	return ta.desc
}

func (ta *Task) GetApps() []string {
	return ta.apps
}

func (ta *Task) GetID() string {
	return ta.id.String()
}

func (ta *Task) GetPrio() int {
	return ta.priority
}

func (ta *Task) GetRemaining() time.Duration {
	return ta.remaining
}

func NewTask(name, desc string, apps []string, priority int, t time.Duration) (*Task, error) {
	if priority >= 5 || priority <= 0 {
		return nil, fmt.Errorf("received error: %w", errInvalidPriority)
	}

	if priority == NotImportantNotUrgent && t != 0 {
		return nil, fmt.Errorf("received error: %w", errPriorityNotMatchingTime)
	}

	if priority != NotImportantNotUrgent && t == 0 {
		return nil, fmt.Errorf("received error: %w", errPriorityNotMatchingTime)
	}

	return &Task{
		id:        uuid.New(),
		name:      name,
		desc:      desc,
		apps:      apps,
		priority:  priority,
		remaining: t,
	}, nil
}
