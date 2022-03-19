//go:build !testing
// +build !testing

package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type gui struct {
	title string

	window fyne.Window

	size fyne.Size
}

func (instance *gui) Render(app fyne.App) {
	instance.window = app.NewWindow(instance.title)

	instance.window.Resize(instance.size)

	data := binding.NewStringList()

	data.Set([]string{
		"hello",
		"foobar",
		"foobar",
		"foobar",
		"foobar",
		"foobar",
		"foobar",
		"foobar",
		"foobar",
		"foobar",
		"foobar",
		"foobar",
		"foobar",
	})

	var list = widget.NewListWithData(data, func() fyne.CanvasObject {
		return widget.NewLabel("template")
	},
		func(i binding.DataItem, o fyne.CanvasObject) {
			o.(*widget.Label).Bind(i.(binding.String))
		})

	var read = widget.NewButton("Read", func() {})

	var export = widget.NewButton("Export", func() {})

	instance.window.SetContent(container.NewBorder(nil, container.New(layout.NewVBoxLayout(), read, export), nil, nil, list))

	instance.window.Show()
}

func ReadSlfEntries() {

}

func WriteSlfEntries() {

}

func Gui() *gui {
	return &gui{
		title: "Slf Exporter",

		size: fyne.NewSize(420, 260),
	}
}
