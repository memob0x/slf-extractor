//go:build gui

package main

import (
	"github.com/memob0x/slf-exporter/gui"
)

func init() {
	global.CreateApp = gui.CreateApp
}
