package integer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestToString(t *testing.T) {
	assert.Equal(t, "123", ToString(123))
	assert.Equal(t, "123", ToString(uint(123)))
	assert.Equal(t, "12", ToString(uint8(12)))
	assert.Equal(t, "123", ToString(uint16(123)))
	assert.Equal(t, "123", ToString(uint32(123)))
	assert.Equal(t, "123", ToString(uint64(123)))
}

func TestToBool(t *testing.T) {
	assert.Equal(t, true, ToBool(12))
	assert.Equal(t, false, ToBool(0))
	assert.Equal(t, false, ToBool(-1))
}

func TestToFloat(t *testing.T) {
	assert.Equal(t, float64(437), ToFloat[int, float64](437))
	assert.Equal(t, float64(1), ToFloat[uint, float64](1))

	assert.Equal(t, float64(16), ToFloat[int8, float64](16))
	assert.Equal(t, float64(1), ToFloat[uint8, float64](1))

	assert.Equal(t, float64(137), ToFloat[int16, float64](137))
	assert.Equal(t, float64(697), ToFloat[uint16, float64](697))

	assert.Equal(t, float64(1), ToFloat[int32, float64](1))
	assert.Equal(t, float64(1), ToFloat[uint32, float64](1))

	assert.Equal(t, float64(1), ToFloat[int64, float64](1))
	assert.Equal(t, float64(1), ToFloat[uint64, float64](1))

	assert.Equal(t, float32(1), ToFloat[int, float32](1))
	assert.Equal(t, float32(1), ToFloat[uint, float32](1))

	assert.Equal(t, float32(1), ToFloat[int8, float32](1))
	assert.Equal(t, float32(1), ToFloat[uint8, float32](1))

	assert.Equal(t, float32(1), ToFloat[int16, float32](1))
	assert.Equal(t, float32(1), ToFloat[uint16, float32](1))

	assert.Equal(t, float32(1), ToFloat[int32, float32](1))
	assert.Equal(t, float32(1), ToFloat[uint32, float32](1))

	assert.Equal(t, float32(1), ToFloat[int64, float32](1))
	assert.Equal(t, float32(1), ToFloat[uint64, float32](1))
}
