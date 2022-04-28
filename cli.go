//go:build cli

package main

import (
	"github.com/memob0x/slf-exporter/cli"
)

func init() {
	global.id = INT_ID_BUILD_CLI

	global.CreateApp = cli.CreateApp
}
