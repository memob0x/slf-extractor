package gui

import (
	"fyne.io/fyne/v2/app"
)

func init() {
	var fyne = app.New()

	var gui = Gui()

	gui.Render(fyne)

	fyne.Run()
}
