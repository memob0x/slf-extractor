//go:build cli

package main

import (
	"github.com/memob0x/slf-extractor/cli"
)

func init() {
	onMain = cli.Launch
}
