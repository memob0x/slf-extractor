package main

import "fyne.io/fyne/v2/app"

func main() {
	var fyne = app.New()

	var gui = Gui()

	gui.Render(fyne)

	fyne.Run()
}
