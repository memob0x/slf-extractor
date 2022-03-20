package utils

import "fmt"

// Gets the given milliseconds formatted in string with "seconds", "minutes" etc...
func FormatMs(milliseconds int64) string {
	if milliseconds < 1000 {
		return fmt.Sprintf("%d milliseconds", milliseconds)
	}

	var seconds int64 = (milliseconds / 1000)

	if seconds < 60 {
		return fmt.Sprintf("%d seconds", seconds)
	}

	var minutes int64 = (milliseconds / (1000 * 60))

	if minutes < 60 {
		return fmt.Sprintf("%d minutes", minutes)
	}

	var hours int64 = (milliseconds / (1000 * 60 * 60))

	if hours < 24 {
		return fmt.Sprintf("%d hours", hours)
	}

	var days int64 = (milliseconds / (1000 * 60 * 60 * 24))

	return fmt.Sprintf("%d hours", days)
}
