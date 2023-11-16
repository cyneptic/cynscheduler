package task

import (
	"time"
)

type TimerService struct {
	Time       time.Duration
	isPaused   bool
	isFinished bool
}

func NewTimerService(t time.Duration) *TimerService {
	return &TimerService{t, false, false}
}

func (s *TimerService) Timer() time.Duration {
	return s.Time
}

func (s *TimerService) Tick() {
	s.Time -= time.Second
	if s.Timer() == 0 {
		s.isFinished = true
		return
	}
}

func (s *TimerService) Done() bool {
	return s.isFinished
}

func (s *TimerService) Toggle() {
	s.isPaused = !s.isPaused
}

func (s *TimerService) Resume() {
	if s.isPaused {
		s.isPaused = false
	}
}

func (s *TimerService) Pause() {
	if !s.isPaused {
		s.isPaused = true
	}
}
