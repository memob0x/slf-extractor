package main

import (
	"fmt"
)

const unit int64 = 1000

const units1stChars string = "kMGTPE"

// Gets the given bytes formatted with unit values.
func FormatBytes(b int64) string {
	if b < unit {
		return fmt.Sprintf("%d %cB", b, units1stChars[0])
	}

	var div int64 = int64(unit)

	var exp int = 0

	for n := b / unit; n >= unit; n /= unit {
		div *= unit

		exp++
	}

	var value float64 = float64(b) / float64(div)

	return fmt.Sprintf("%.1f %cB", value, units1stChars[exp])
}
