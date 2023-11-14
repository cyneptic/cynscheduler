package task_test

import (
	"testing"
	"time"

	"github.com/cyneptic/cynscheduler/task"
	"github.com/stretchr/testify/assert"
)

func TestModel(t *testing.T) {
	testCases := []struct {
		desc        string
		name        string
		description string
		remaining   int
	}{
		{
			desc:        "Valid Model",
			name:        "Walk",
			description: "Go for a walk",
			remaining:   1,
		},
	}
	for _, tC := range testCases {
		tc := tC
		t.Run(tC.desc, func(t *testing.T) {
			actual := task.NewTask(tc.name, tc.description, tc.remaining)
			assert.Equal(t, tc.name, actual.Name)
			assert.Equal(t, tc.description, actual.Description)
			assert.Equal(t, time.Duration(tc.remaining)*time.Minute, actual.Remaining)
		})
	}
}
