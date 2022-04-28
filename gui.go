//go:build gui

package main

import (
	"github.com/memob0x/slf-exporter/gui"
)

func init() {
	global.id = INT_ID_BUILD_GUI

	global.CreateApp = gui.CreateApp
}
