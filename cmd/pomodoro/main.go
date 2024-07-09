package main

import (
	"fmt"
	"github.com/rivo/tview"
	"os"
	"pomodoro/internal/sound"
	"pomodoro/internal/timer"
	"pomodoro/internal/ui"
)

var (
	app         *tview.Application
	currentPage string
	basePath    string
)

const (
	pageCount      = 2
	KEY_SOUND_PATH = "/assets/wav/key.wav"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {

	timerInstance := timer.GetTimerInstance()

	homePath, err := os.UserHomeDir()
	check(err)

	basePath = fmt.Sprintf("%s/.config/go-pomo", homePath)
    _, err = os.ReadDir(basePath)

    if err != nil {
        fmt.Println("Please initialize go-pomo using the install.sh script first.")
        return
    }

	audioPlayer := sound.GetPlayerInstance(basePath + KEY_SOUND_PATH)

	app = tview.NewApplication()

	currentPage = "Landing"
	pages := ui.SetupPages(app, &currentPage, timerInstance, audioPlayer)

	if err := app.SetRoot(pages, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}

}
