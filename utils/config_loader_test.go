package utils_test

import (
	"testing"

	"github.com/cyneptic/cynscheduler/utils"
	"github.com/stretchr/testify/require"
)

func TestConfigLoader(t *testing.T) {
	_, err := utils.Config_loader()
	require.NoError(t, err)
}

func TestSampleConfig(t *testing.T) {
	result, err := utils.Config_loader()
	require.NoError(t, err)

	require.Equal(t, "Walk", result[0].Name)
}
