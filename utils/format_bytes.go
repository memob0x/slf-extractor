package utils

import (
	"fmt"
)

// Gets the given bytes formatted with unit values.
func FormatBytes(b int64) string {
	const unitSize int64 = 1000

	const unitsLeadingChars string = "kMGTPE"

	if b < unitSize {
		return fmt.Sprintf("%d %cB", b, unitsLeadingChars[0])
	}

	var div int64 = int64(unitSize)

	var exp int = 0

	for n := b / unitSize; n >= unitSize; n /= unitSize {
		div *= unitSize

		exp++
	}

	var value float64 = float64(b) / float64(div)

	return fmt.Sprintf("%.1f %cB", value, unitsLeadingChars[exp])
}
