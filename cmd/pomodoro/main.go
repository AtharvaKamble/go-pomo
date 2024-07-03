package main

import (
	"github.com/rivo/tview"
	"pomodoro/internal/models"
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

	timer := models.GetTimerInstance()

	delta := make(chan time.Duration)

	total := 1500 * time.Second

	app = tview.NewApplication()
	// box := tview.NewBox().SetBorder(true).SetTitle("GoPomo")

	// there must be global context to store which page was selected

	currentPage = "Landing"
	pages := ui.SetupPages(app, &currentPage)

	textView = tview.NewTextView().
		SetDynamicColors(true).
		SetRegions(true).
		SetChangedFunc(func() {
			app.Draw()
		})

	go func() {
		for t := range delta {
			curTime := total - t.Truncate(time.Second)

			app.QueueUpdateDraw(func() {
				textView.SetText(curTime.String())
			})
		}

	}()

	go timer.Tick(total, delta)

	if err := app.SetRoot(pages, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}

}
