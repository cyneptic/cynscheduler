package task_test

import (
	"testing"
	"time"

	"github.com/cyneptic/cynscheduler/task"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTaskCreation(t *testing.T) {
	t1, err := task.NewTask("Walk", "Go for a walk", []string{}, task.ImportantAndUrgent, 1*time.Hour)

	require.NoError(t, err)
	assert.NotNil(t, t1)

	assert.Equal(t, "Walk", t1.GetName())
	assert.Equal(t, "Go for a walk", t1.GetDesc())
	assert.Equal(t, []string{}, t1.GetApps())
	assert.Equal(t, 1, t1.GetPrio())
	assert.Equal(t, 1*time.Hour, t1.GetRemaining())
}

func TestTaskCreationInvalidPriority(t *testing.T) {
	t1, err := task.NewTask("Walk", "Go for a walk", []string{}, 5, 1*time.Hour)
	require.Error(t, err)
	assert.Nil(t, t1)
}

func TestTaskCreationInvalidPriorityToTime(t *testing.T) {
	t1, err := task.NewTask("Walk", "Go for a walk", []string{}, 4, 1*time.Hour)
	require.Error(t, err)
	assert.Nil(t, t1)

	t2, err := task.NewTask("Walk", "Go for a walk", []string{}, 1, 0)
	require.Error(t, err)
	assert.Nil(t, t2)
}
