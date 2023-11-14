package utils_test

import (
	"testing"
	"time"

	"github.com/cyneptic/cynscheduler/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestConfigLoader(t *testing.T) {
	assert := assert.New(t)
	tasks, err := utils.LoadConfig("sample_config.json")
	require.NoError(t, err)

	assert.Equal("walk", tasks[0].Name)
	assert.Equal("go for a walk", tasks[0].Description)
	assert.Equal(5*time.Minute, tasks[0].Timer.Timer())
	assert.Equal("groceries", tasks[1].Name)
	assert.Equal("buy groceries", tasks[1].Description)
	assert.Equal(5*time.Minute, tasks[1].Timer.Timer())
}
