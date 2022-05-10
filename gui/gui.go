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

const LABEL_BUTTON_FILE_BROWSER_DEFAULT = "(None selected)"
const LABEL_TITLE_MAIN_WINDOW = "Slf Exporter"

type Store struct {
	App fyne.App

	SlfPath string

	DestinationPath string

	Entries binding.StringList

	List *widget.List

	ButtonSlfFile *widget.Button

	ButtonDestinationPath *widget.Button

	ButtonExport *widget.Button
}

func (instance *Store) GetMainContent() *fyne.Container {
	return container.NewBorder(
		container.New(
			layout.NewAdaptiveGridLayout(3),

			container.New(
				layout.NewVBoxLayout(),

				widget.NewLabel("Slf file:"),

				instance.ButtonSlfFile,
			),

			container.New(
				layout.NewVBoxLayout(),

				widget.NewLabel("Extraction folder:"),

				instance.ButtonDestinationPath,
			),

			instance.ButtonExport,
		),

		nil,

		nil,

		nil,

		instance.List,
	)
}

func (instance *Store) Render(app fyne.App) {
	instance.App = app

	var window fyne.Window = app.NewWindow(LABEL_TITLE_MAIN_WINDOW)

	window.Resize(fyne.NewSize(420, 260))

	window.SetContent(instance.GetMainContent())

	instance.Refresh()

	window.Show()
}

type BrowserDialogType int64

const (
	BrowserDialogTypeFolder BrowserDialogType = iota

	BrowserDialogypeSlf
)

func (instance *Store) RenderFileBrowserDialog(dialogType BrowserDialogType) {
	var window fyne.Window = instance.App.NewWindow(LABEL_TITLE_MAIN_WINDOW + ": file browser")

	window.Resize(fyne.NewSize(720, 576))

	var fileChooser *dialog.FileDialog

	var onSelection = func() {
		window.Close()

		instance.Refresh()
	}

	switch dialogType {

	case BrowserDialogTypeFolder:
		fileChooser = dialog.NewFolderOpen(func(selection fyne.ListableURI, err error) {
			if selection != nil {
				instance.DestinationPath = selection.Path()
			}

			onSelection()
		}, window)

		break

	case BrowserDialogypeSlf:
		fileChooser = dialog.NewFileOpen(func(selection fyne.URIReadCloser, err error) {
			if selection != nil {
				instance.SlfPath = selection.URI().Path()
			}

			onSelection()
		}, window)

		fileChooser.SetFilter(storage.NewExtensionFileFilter([]string{".slf"}))

		break

	default:
		return
	}

	window.Show()

	fileChooser.Show()
}

func IsValidString(str string) bool {
	return len(strings.Trim(str, " ")) > 0
}

func (instance *Store) HasSlfPath() bool {
	return IsValidString(instance.SlfPath)
}

func (instance *Store) HasDestinationPath() bool {
	return IsValidString(instance.DestinationPath)
}

func (instance *Store) Export() {
	if instance.Entries.Length() <= 0 {
		return
	}

	var _, _, _, err = utils.ExtractSlfEntries(
		instance.SlfPath,

		instance.DestinationPath,

		8,

		func(stats fs.FileInfo) {},

		func(perc float64) {},

		func(header utils.SlfHeader) {},

		func(file *os.File) {},

		func(files []*os.File) {},
	)

	if err != nil {
		// TODO: error message

		return
	}

	// TODO: success message

	instance.SlfPath = ""

	instance.Refresh()
}

func (instance *Store) HasEntries() bool {
	return instance.Entries.Length() > 0
}

func (instance *Store) UpdateButtonSlfFileLabel() {
	instance.ButtonSlfFile.Text = LABEL_BUTTON_FILE_BROWSER_DEFAULT

	if instance.HasSlfPath() {
		instance.ButtonSlfFile.Text = instance.SlfPath
	}
}

func (instance *Store) UpdateButtonDestinationPathLabel() {
	instance.ButtonDestinationPath.Text = LABEL_BUTTON_FILE_BROWSER_DEFAULT

	if !instance.HasDestinationPath() {
		return
	}

	var label string = instance.DestinationPath

	if label == "." {
		label = "(Current folder)"
	}

	instance.ButtonDestinationPath.Text = label
}

func (instance *Store) CanExport() bool {
	return instance.HasSlfPath() && instance.HasDestinationPath() && instance.HasEntries()
}

func (instance *Store) UpdateButtonExportState() {
	instance.ButtonExport.Disable()

	if instance.CanExport() {
		instance.ButtonExport.Enable()
	}
}

func (instance *Store) RefreshButtons() error {
	instance.UpdateButtonSlfFileLabel()

	instance.UpdateButtonDestinationPathLabel()

	instance.UpdateButtonExportState()

	instance.ButtonSlfFile.Refresh()

	instance.ButtonDestinationPath.Refresh()

	instance.ButtonExport.Refresh()

	return nil
}

func (instance *Store) RefreshEntries() error {
	instance.Entries.Set([]string{})

	if !instance.HasSlfPath() {
		return nil
	}

	var entries, _, _, err = utils.ReadSlfFile(
		instance.SlfPath,

		8,

		func(stats fs.FileInfo) {},

		func(perc float64) {},

		func(header utils.SlfHeader) {},
	)

	if err != nil {
		// TODO: display error message

		return err
	}

	for i, j := 0, len(entries); i < j; i++ {
		var entry utils.SlfEntry = entries[i]

		instance.Entries.Append(entry.Name)
	}

	return nil
}

func (instance *Store) Refresh() error {
	var err error = instance.RefreshEntries()

	if err != nil {
		return err
	}

	err = instance.RefreshButtons()

	if err != nil {
		return err
	}

	return nil
}

func New() *Store {
	var entries = binding.NewStringList()

	var instance = &Store{
		DestinationPath: ".",

		Entries: entries,

		List: widget.NewListWithData(
			entries,

			func() fyne.CanvasObject {
				return widget.NewLabel("")
			},

			func(i binding.DataItem, o fyne.CanvasObject) {
				o.(*widget.Label).Bind(i.(binding.String))
			},
		),

		ButtonSlfFile: widget.NewButton("", nil),

		ButtonDestinationPath: widget.NewButton("", nil),

		ButtonExport: widget.NewButton("Export", nil),
	}

	instance.ButtonSlfFile.OnTapped = func() {
		// TODO: hide displayed error messages

		instance.RenderFileBrowserDialog(BrowserDialogypeSlf)
	}

	instance.ButtonDestinationPath.OnTapped = func() {
		instance.RenderFileBrowserDialog(BrowserDialogTypeFolder)
	}

	instance.ButtonExport.OnTapped = instance.Export

	return instance
}

func Launch() {
	var fyne = app.New()

	var gui = New()

	gui.Render(fyne)

	fyne.Run()
}
