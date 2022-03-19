package main

import (
	"fyne.io/fyne/v2/app"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	var fyne = app.New()

	var gui = Gui()

	gui.Render(fyne)

	fyne.Run()
}
