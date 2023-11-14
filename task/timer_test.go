package task_test

import (
	"testing"
	"time"

	"github.com/cyneptic/cynscheduler/task"
	"github.com/stretchr/testify/assert"
)

func TestTimer(t *testing.T) {
	timer := task.NewTimerService(10 * time.Second)

	timer.Tick()
	assert.Equal(t, time.Second*9, timer.Timer())
	timer.Tick()
	assert.Equal(t, time.Second*8, timer.Timer())
}
