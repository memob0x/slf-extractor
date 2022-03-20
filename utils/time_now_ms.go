package utils

import "time"

// Returns the current time in ms
func TimeNowMs() int64 {
	return int64(time.Nanosecond) * time.Now().UnixNano() / int64(time.Millisecond)
}
