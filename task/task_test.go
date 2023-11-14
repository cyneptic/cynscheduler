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
		important   bool
		urgent      bool
		remaining   int
	}{
		{
			desc:        "Important and Not Urgent Model",
			name:        "Walk",
			description: "Go for a walk",
			important:   true,
			urgent:      false,
			remaining:   1,
		},
		{
			desc:        "Important and Urgent Model",
			name:        "Walk",
			description: "Go for a walk",
			important:   true,
			urgent:      true,
			remaining:   1,
		},
		{
			desc:        "Not Important but Urgent Model",
			name:        "Walk",
			description: "Go for a walk",
			important:   false,
			urgent:      true,
			remaining:   1,
		},
		{
			desc:        "Not Important and Not Urgent Model",
			name:        "Walk",
			description: "Go for a walk",
			important:   false,
			urgent:      false,
			remaining:   1,
		},
	}

	for _, tC := range testCases {
		tc := tC
		t.Run(tC.desc, func(t *testing.T) {
			actual := task.NewTask(tc.name, tc.description, tc.remaining, tc.important, tc.urgent)
			assert.Equal(t, tc.name, actual.Name)
			assert.Equal(t, tc.description, actual.Description)
			assert.Equal(t, time.Duration(tc.remaining)*time.Minute, actual.Timer.Timer())
			assert.Equal(t, tc.important, actual.Important)
			assert.Equal(t, tc.urgent, actual.Urgent)
		})
	}
}
