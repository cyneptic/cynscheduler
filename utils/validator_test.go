package utils_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/cyneptic/cynscheduler/task"
	"github.com/cyneptic/cynscheduler/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestValidationNameAndDesc(t *testing.T) {
	t.Run("Valid Entries", func(t *testing.T) {
		t.Parallel()
		err := utils.ValidateNameAndDescription("Walk", "Go for a walk")
		require.NoError(t, err)
	})

	t.Run("Invalid Name", func(t *testing.T) {
		t.Parallel()
		err := utils.ValidateNameAndDescription("", "Go for a walk")
		require.Error(t, err)
	})

	t.Run("Invalid Desc", func(t *testing.T) {
		t.Parallel()
		err := utils.ValidateNameAndDescription("Walk", "")
		require.Error(t, err)
	})
}

func TestValidatePriority(t *testing.T) {
	testCases := []struct {
		desc             string
		priority         string
		expectedPriority int
		expectErr        bool
	}{
		{
			desc:             "priority valid #1",
			priority:         fmt.Sprintf("%d", task.ImportantAndUrgent),
			expectedPriority: 1,
			expectErr:        false,
		},
		{
			desc:             "priority valid #2",
			priority:         fmt.Sprintf("%d", task.NotImportantUrgent),
			expectedPriority: 2,
			expectErr:        false,
		},
		{
			desc:             "priority valid #3",
			priority:         fmt.Sprintf("%d", task.ImportantNotUrgent),
			expectedPriority: 3,
			expectErr:        false,
		},
		{
			desc:             "priority valid #4",
			priority:         fmt.Sprintf("%d", task.NotImportantNotUrgent),
			expectedPriority: 4,
			expectErr:        false,
		},
		{
			desc:             "priority invalid",
			priority:         "5",
			expectedPriority: -1,
			expectErr:        true,
		},
	}

	for _, tC := range testCases {
		tC := tC
		t.Run(tC.desc, func(t *testing.T) {
			t.Parallel()
			actual, err := utils.ValidatePriority(tC.priority)
			if tC.expectErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}

			assert.Equal(t, tC.expectedPriority, actual)
		})
	}
}

func TestValidateDuration(t *testing.T) {
	testCases := []struct {
		desc      string
		inputStr  string
		expected  time.Duration
		expectErr bool
	}{
		{
			desc:      "Test 5 Minutes",
			inputStr:  "5",
			expected:  5 * time.Minute,
			expectErr: false,
		},
		{
			desc:      "Test 1 Hour",
			inputStr:  "60",
			expected:  1 * time.Hour,
			expectErr: false,
		},
		{
			desc:      "Zero Duration",
			inputStr:  "0",
			expected:  0,
			expectErr: false,
		},
		{
			desc:      "Negative Duration",
			inputStr:  "-5",
			expected:  -1,
			expectErr: true,
		},
		{
			desc:      "Invalid String",
			inputStr:  "",
			expected:  -1,
			expectErr: true,
		},
	}
	for _, tC := range testCases {
		tc := tC
		t.Run(tc.desc, func(t *testing.T) {
			ti, err := utils.ValidateDuration(tc.inputStr)

			if tc.expectErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}

			assert.Equal(t, tc.expected, ti)
		})
	}
}
