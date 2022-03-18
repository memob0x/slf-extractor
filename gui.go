// +build !testing

package main

import "fyne.io/fyne/v2"

type gui struct {
	title string

	window fyne.Window

	size fyne.Size
}

func (instance *gui) Render(app fyne.App) {
	instance.window = app.NewWindow(instance.title)

	instance.window.Resize(instance.size)

	instance.window.Show()
}

func Gui() *gui {
	return &gui{
		title: "Slf Exporter",

		size: fyne.NewSize(420, 260),
	}
}
