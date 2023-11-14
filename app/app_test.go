package app_test

import (
	"testing"

	"github.com/cyneptic/cynscheduler/app"
	"github.com/cyneptic/cynscheduler/task"
	"github.com/stretchr/testify/assert"
)

func TestSortByMatrix(t *testing.T) {
	testCases := []struct {
		desc  string
		tasks []struct {
			name      string
			important bool
			urgent    bool
		}
		expected []struct {
			name      string
			important bool
			urgent    bool
		}
	}{
		{
			desc: "important and urgent",
			tasks: []struct {
				name      string
				important bool
				urgent    bool
			}{
				{"1", true, true},
				{"2", true, true},
				{"3", true, false},
				{"4", true, false},
				{"7", false, false},
				{"5", true, false},
				{"6", false, true},
				{"8", false, false},
				{"before3", true, true},
			},
			expected: []struct {
				name      string
				important bool
				urgent    bool
			}{
				{"1", true, true},
				{"2", true, true},
				{"before3", true, true},
				{"3", true, false},
				{"4", true, false},
				{"5", true, false},
				{"6", false, true},
				{"7", false, false},
				{"8", false, false},
			},
		},
	}
	for _, tC := range testCases {
		tc := tC
		t.Run(tC.desc, func(t *testing.T) {
			q := []*task.Task{}
			for _, v := range tc.tasks {
				q = append(q, task.NewTask(v.name, "", 0, v.important, v.urgent))
			}

			app := app.NewApp(q, 1)
			assert.Equal(t, tc.expected[0].name, app.Tasks[0].Name)
			assert.Equal(t, tc.expected[1].name, app.Tasks[1].Name)
			assert.Equal(t, tc.expected[2].name, app.Tasks[2].Name)
			assert.Equal(t, tc.expected[3].name, app.Tasks[3].Name)
			assert.Equal(t, tc.expected[4].name, app.Tasks[4].Name)
			assert.Equal(t, tc.expected[5].name, app.Tasks[5].Name)
			assert.Equal(t, tc.expected[6].name, app.Tasks[6].Name)
			assert.Equal(t, tc.expected[7].name, app.Tasks[7].Name)
			assert.Equal(t, tc.expected[8].name, app.Tasks[8].Name)
		})
	}
}
