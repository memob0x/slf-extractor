package gui

import (
	"io/fs"
	"os"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
	"github.com/memob0x/slf-extractor/utils"
)

type GuiStore struct {
	App fyne.App

	Path string

	Entries binding.StringList

	List *widget.List

	ButtonReadEntries *widget.Button

	ButtonWriteEntries *widget.Button
}

func (instance *GuiStore) Render(app fyne.App) {
	instance.App = app

	var window fyne.Window = app.NewWindow("Slf Exporter")

	window.Resize(fyne.NewSize(420, 260))

	var windowContent *fyne.Container = container.NewBorder(
		container.New(
			layout.NewVBoxLayout(),

			instance.ButtonReadEntries,

			instance.ButtonWriteEntries,
		),

		nil,

		nil,

		nil,

		instance.List,
	)

	instance.Refresh()

	window.SetContent(windowContent)

	window.Show()
}

func (instance *GuiStore) RenderSlfFileChooser() {
	var window fyne.Window = instance.App.NewWindow("Slf Exporter: file chooser")

	window.Resize(fyne.NewSize(420, 260))

	var fileChooser *dialog.FileDialog = dialog.NewFileOpen(func(file fyne.URIReadCloser, err error) {
		instance.Path = file.URI().Path()

		window.Close()

		instance.Refresh()
	}, window)

	fileChooser.SetFilter(storage.NewExtensionFileFilter([]string{".slf"}))

	window.Show()

	fileChooser.Show()
}

func (instance *GuiStore) HasPath() bool {
	return len(strings.Trim(instance.Path, " ")) > 0
}

func (instance *GuiStore) Export() {
	if instance.Entries.Length() <= 0 {
		return
	}

	var _, _, _, err = utils.ExtractSlfEntries(instance.Path, ".", 8, func(stats fs.FileInfo) {}, func(perc float64) {}, func(header utils.SlfHeader) {}, func(file *os.File) {}, func(files []*os.File) {})

	if err != nil {
		// TODO: error message

		return
	}

	// TODO: success message

	instance.Path = ""

	instance.Refresh()
}

func (instance *GuiStore) Refresh() error {
	instance.Entries.Set([]string{})

	instance.ButtonWriteEntries.Disable()

	if !instance.HasPath() {
		return nil
	}

	var entries, _, _, err = utils.ReadSlfFile(instance.Path, 8, func(stats fs.FileInfo) {}, func(perc float64) {}, func(header utils.SlfHeader) {})

	if err != nil {
		// TODO: display errors

		return err
	}

	for i, j := 0, len(entries); i < j; i++ {
		var entry utils.SlfEntry = entries[i]

		instance.Entries.Append(entry.Name)
	}

	instance.ButtonWriteEntries.Enable()

	return nil
}

func CreateGui() *GuiStore {
	var entries = binding.NewStringList()

	var instance = &GuiStore{
		Entries: entries,

		List: widget.NewListWithData(
			entries,

			func() fyne.CanvasObject {
				return widget.NewLabel("template")
			},

			func(i binding.DataItem, o fyne.CanvasObject) {
				o.(*widget.Label).Bind(i.(binding.String))
			},
		),

		ButtonReadEntries: widget.NewButton("Open", nil),

		ButtonWriteEntries: widget.NewButton("Export", nil),
	}

	instance.ButtonWriteEntries.OnTapped = instance.Export

	instance.ButtonReadEntries.OnTapped = func() {
		// TODO: hide displayed errors

		instance.RenderSlfFileChooser()
	}

	return instance
}

func CreateApp() {
	var fyne = app.New()

	var gui = CreateGui()

	gui.Render(fyne)

	fyne.Run()
}
