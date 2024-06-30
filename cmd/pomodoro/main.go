package main

import (
	"fmt"
	"pomodoro/internal/models"
	"time"
)

func main() {

	timer := models.GetTimerInstance()

	delta := make(chan time.Duration)
	go timer.Tick(10*time.Second, delta)

	for t := range delta {
		fmt.Println(t)
	}
}
