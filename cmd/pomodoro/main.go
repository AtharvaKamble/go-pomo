package main

import (
	"github.com/rivo/tview"
	"pomodoro/internal/timer"
	"pomodoro/internal/ui"
	"time"
)

var (
	app         *tview.Application
	textView    *tview.TextView
	currentPage string
)

const pageCount = 2

func main() {

	timer := timer.GetTimerInstance()

	delta := make(chan time.Duration)

	//total := 1500 * time.Second

	app = tview.NewApplication()
	// box := tview.NewBox().SetBorder(true).SetTitle("GoPomo")

	// there must be global context to store which page was selected

	currentPage = "Landing"
	pages := ui.SetupPages(app, &currentPage, timer, delta)

	textView = tview.NewTextView().
		SetDynamicColors(true).
		SetRegions(true).
		SetChangedFunc(func() {
			app.Draw()
		})

	if err := app.SetRoot(pages, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}

}
