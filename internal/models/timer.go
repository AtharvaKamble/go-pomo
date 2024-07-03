package models

import (
	"sync"
	"time"
)

type Timer struct {
	Start time.Time
}

var (
	cache     *Timer
	cacheOnce sync.Once
)

// GetTimerInstance creates a single new instance utilizing the Singleton pattern
func GetTimerInstance() *Timer {
	// Do only run once on every sync.Once instance
	cacheOnce.Do(func() {
		cache = &Timer{Start: time.Now()}
	})

	return cache
}

// GetDelta provides the time elapsed from the creation of the timer *Timer struct instance
func (timer *Timer) GetDelta() time.Duration {
	return time.Since(timer.Start)
}

// Tick is a GoRoutine that ticks every second for a given duration
func (timer *Timer) Tick(duration time.Duration, tickChannel chan time.Duration) {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {

		select {
		case <-ticker.C:
			tickChannel <- timer.GetDelta()

			if timer.GetDelta() >= duration {
				close(tickChannel)
			}
		}

	}
}
