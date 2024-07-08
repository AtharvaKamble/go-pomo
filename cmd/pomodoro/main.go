package main

import (
	"github.com/rivo/tview"
	"os"
	"pomodoro/internal/sound"
	"pomodoro/internal/timer"
	"pomodoro/internal/ui"
)

var (
	app         *tview.Application
	currentPage string
)

const pageCount = 2

func main() {

	timerInstance := timer.GetTimerInstance()

	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	audioPlayer := sound.GetPlayerInstance(path + "/../../assets/wav/key.wav")

	app = tview.NewApplication()

	currentPage = "Landing"
	pages := ui.SetupPages(app, &currentPage, timerInstance, audioPlayer)

	if err := app.SetRoot(pages, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}

}
