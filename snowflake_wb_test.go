package snowflake

import (
	"strconv"
	"testing"

	"gotest.tools/assert"
)

func TestCalculateMaxForMaskNegative(t *testing.T) {
	assert.Equal(t, calculateMaxForMask(-199), int64(-1))
	assert.Equal(t, calculateMaxForMask(0), int64(-1))
}

func TestCalculateMaxForMaskOverflow(t *testing.T) {
	assert.Equal(t, calculateMaxForMask(65), int64(-1))
	assert.Equal(t, calculateMaxForMask(64), int64(-1))
}

func TestCalculateMaxForMask(t *testing.T) {
	s := "1"
	for i := 1; i < 64; i++ {
		val, err := strconv.ParseInt(s, 2, 64)
		assert.NilError(t, err)
		assert.Equal(t, calculateMaxForMask(i), int64(val))
		s += "1"
	}
}
