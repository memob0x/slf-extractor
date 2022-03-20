package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTimeNowMs(t *testing.T) {
	assert.IsTypef(t, int64(0), TimeNowMs(), "should return an int64")
}
