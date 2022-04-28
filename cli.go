//go:build cli

package main

import (
	"github.com/memob0x/slf-exporter/cli"
)

func init() {
	global.CreateApp = cli.CreateApp
}
