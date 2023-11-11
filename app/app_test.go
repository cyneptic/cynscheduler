package app_test

import (
	"testing"
	"time"

	"github.com/cyneptic/cynscheduler/app"
	"github.com/cyneptic/cynscheduler/task"
	"github.com/stretchr/testify/assert"
)

var (
	sampleRemainingDuration = 8 * time.Hour
	sampleWaitMargin        = 100 * time.Millisecond
)

func generateSampleApp() *app.App {
	return app.NewApp(make(map[string]*task.Task), sampleRemainingDuration)
}

func TestNewApp(t *testing.T) {
	a := generateSampleApp()
	assert.NotNil(t, a)
}

func TestAppStart(t *testing.T) {
	a := generateSampleApp()

	assert.Equal(t, sampleRemainingDuration, a.RemainingTime)
	go a.Start()
	time.Sleep(sampleWaitMargin)
	assert.NotEqual(t, sampleRemainingDuration, a.RemainingTime)
}

func TestAppPauseAndResume(t *testing.T) {
	a := generateSampleApp()

	assert.Equal(t, sampleRemainingDuration, a.RemainingTime)
	go a.Start()

	time.Sleep(sampleWaitMargin)
	assert.NotEqual(t, sampleRemainingDuration, a.RemainingTime)

	a.Pause()
	before := a.RemainingTime
	time.Sleep(sampleWaitMargin)
	after := a.RemainingTime

	assert.Equal(t, before, after)

	a.Resume()
	time.Sleep(sampleWaitMargin)
	after = a.RemainingTime
	assert.NotEqual(t, before, after)
}
