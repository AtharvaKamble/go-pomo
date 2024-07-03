package ui

import (
	"fmt"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func SetupPages(app *tview.Application, currentPage *string) *tview.Pages {

	pages := tview.NewPages()

	// Page 1 - landing page
	pages.AddPage("Landing",

		tview.NewFlex().SetDirection(tview.FlexRow).AddItem(
			tview.NewModal().
				SetBackgroundColor(tcell.Color148).
				SetTextColor(tcell.Color16).
				SetText(fmt.Sprintf("Select a duration for your pomodoro")).
				AddButtons([]string{"15m", "25m", "45m"}).
				SetDoneFunc(func(buttonIndex int, buttonLabel string) {
					if buttonLabel == "15m" {
						pages.SwitchToPage("Timer")
						*currentPage = "Timer"
					} else if buttonLabel == "25m" {
						pages.SwitchToPage("Timer")
						*currentPage = "Timer"
					} else if buttonLabel == "45m" {
						pages.SwitchToPage("Timer")
						*currentPage = "Timer"
					} else {
						app.Stop()
					}
					fmt.Println(*currentPage)
				}), 0, 3, true),
		false,
		true)

	// Page 2 - timer page
	pages.AddPage("Timer",

		tview.NewModal().
			SetBackgroundColor(tcell.Color148).
			SetTextColor(tcell.Color16).
			SetText(fmt.Sprintf("This is the timer page")).
			AddButtons([]string{"Back", "Quit"}).
			SetDoneFunc(func(buttonIndex int, buttonLabel string) {
				if buttonLabel == "Back" {
					pages.SwitchToPage("Landing")
					*currentPage = "Landing"
					fmt.Println(*currentPage)
				} else if buttonLabel == "Quit" {
					app.Stop()
				} else {
					app.Stop()
				}
			}),
		false,
		false)

	return pages
}
