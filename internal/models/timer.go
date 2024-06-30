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

func GetTimerInstance() *Timer {
	cacheOnce.Do(func() {
		cache = &Timer{Start: time.Now()}
	})

	return cache
}

func (timer *Timer) GetDelta() time.Duration {
	return time.Since(timer.Start)
}

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
