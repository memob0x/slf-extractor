package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
)

type gui struct {
	app fyne.App

	entries binding.StringList

	list *widget.List

	buttonReadEntries *widget.Button

	buttonWriteEntries *widget.Button
}

func (instance *gui) Render(app fyne.App) {
	instance.app = app

	var window fyne.Window = app.NewWindow("Slf Exporter")

	window.Resize(fyne.NewSize(420, 260))

	var windowContent *fyne.Container = container.NewBorder(
		container.New(
			layout.NewVBoxLayout(),

			instance.buttonReadEntries,

			instance.buttonWriteEntries,
		),

		nil,

		nil,

		nil,

		instance.list,
	)

	instance.Refresh()

	window.SetContent(windowContent)

	window.Show()
}

func (instance *gui) WriteSlfFileEntriesFiles() {
	// TODO: actual writing of entries
	// TODO: clear/reset gui

	instance.Refresh()
}

func (instance *gui) RenderSlfFileChooser() {
	var window fyne.Window = instance.app.NewWindow("Slf Exporter: file chooser")

	window.Resize(fyne.NewSize(420, 260))

	var fileChooser *dialog.FileDialog = dialog.NewFileOpen(func(fyne.URIReadCloser, error) {

	}, window)

	fileChooser.SetFilter(storage.NewExtensionFileFilter([]string{".slf"}))

	window.Show()

	fileChooser.Show()
}

func (instance *gui) Refresh() {
	// if file ok
	// TODO: display list
	// TODO: enable "write" button

	// if file not ok
	// TODO: display error with "x" button

	instance.entries.Set([]string{"hello"})

	instance.buttonReadEntries.OnTapped = func() {
		// TODO: hide former errors

		instance.RenderSlfFileChooser()
	}

	instance.buttonReadEntries.Refresh()
}

func Gui() *gui {
	var entries = binding.NewStringList()

	return &gui{
		entries: entries,

		list: widget.NewListWithData(
			entries,

			func() fyne.CanvasObject {
				return widget.NewLabel("template")
			},

			func(i binding.DataItem, o fyne.CanvasObject) {
				o.(*widget.Label).Bind(i.(binding.String))
			},
		),

		buttonReadEntries: widget.NewButton("Read", nil),

		buttonWriteEntries: widget.NewButton("Export", nil),
	}
}
