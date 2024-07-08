package ui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"pomodoro/internal/sound"
	"pomodoro/internal/timer"
	"time"
)

var (
	buttonsModal     *tview.Modal
	pomodoroDuration time.Duration
)

func SetupPages(app *tview.Application, currentPage *string, timer *timer.Timer, ap *sound.AudioPlayer) *tview.Pages {
	delta := make(chan time.Duration)

	pages := tview.NewPages()

	modal1 := tview.NewModal().
		SetBackgroundColor(tcell.Color148).
		SetTextColor(tcell.Color16).
		SetText("Select a duration for your pomodoro").
		AddButtons([]string{"15m", "25m", "45m"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {

			if buttonLabel == "15m" {
				pomodoroDuration = 15 * 60
			} else if buttonLabel == "25m" {
				pomodoroDuration = 25 * 60
			} else if buttonLabel == "45m" {
				pomodoroDuration = 45 * 60
			} else {
				app.Stop()
			}

			pages.SwitchToPage("Timer")
			*currentPage = "Timer"
			//fmt.Println(pomodoroDuration)
			go timer.Tick(pomodoroDuration*time.Second, delta)

			go updateUI(delta, app, buttonsModal, pomodoroDuration, ap)
		})

	page1 := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(tview.NewBox(), 0, 1, false).
		AddItem(tview.NewFlex().
			AddItem(tview.NewBox(), 0, 1, false).
			AddItem(modal1, 0, 5, true).
			AddItem(tview.NewBox(), 0, 1, false), 0, 2, true).
		AddItem(tview.NewBox(), 0, 1, false).
		AddItem(tview.NewFlex().
			AddItem(tview.NewBox(), 0, 1, false).
			AddItem(tview.NewBox(), 0, 2, true).
			AddItem(tview.NewBox(), 0, 1, false), 0, 2, true).
		AddItem(tview.NewBox(), 0, 1, false)

	// Page 2 - timer page
	buttonsModal = tview.NewModal().
		SetBackgroundColor(tcell.Color148).
		SetTextColor(tcell.Color16).
		AddButtons([]string{"Reset", "Quit"}).
		SetText("Starting pomodoro session, get ready...").
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			if buttonLabel == "Reset" {
				pages.SwitchToPage("Landing")
				*currentPage = "Landing"
			} else if buttonLabel == "Quit" {
				app.Stop()
			} else {
				app.Stop()
			}
		})

	//timerModal := tview.NewBox().SetBorder(true).SetTitle("Timer box")

	page2 := tview.NewFlex().SetDirection(tview.FlexRow).AddItem(buttonsModal, 0, 2, false).AddItem(buttonsModal, 0, 2, true)

	pages.AddPage("Landing", page1, true, true)
	pages.AddPage("Timer", page2, false, false)

	return pages
}

func updateUI(delta chan time.Duration, app *tview.Application, buttonsModal *tview.Modal, pomodoroDuration time.Duration, ap *sound.AudioPlayer) {

	for t := range delta {
		elapsedTime := pomodoroDuration*time.Second - t.Truncate(time.Second)
		ap.PlaySound()

		app.QueueUpdateDraw(func() {
			buttonsModal.SetText(elapsedTime.String())
		})
	}

}
