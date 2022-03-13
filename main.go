package main

import "fyne.io/fyne/v2/app"

func main() {
	app := app.New()

	c := getGui()

	c.render(app)

	app.Run()
}
