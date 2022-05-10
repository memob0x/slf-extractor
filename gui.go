//go:build gui

package main

import (
	"github.com/memob0x/slf-extractor/gui"
)

func init() {
	onMain = gui.Launch
}
