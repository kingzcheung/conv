package boolean

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestToString(t *testing.T) {
	assert.Equal(t, "true", ToString(true))
	assert.Equal(t, "false", ToString(false))
	assert.Equal(t, "1", ToString(true, NumberString))
	assert.Equal(t, "0", ToString(false, NumberString))
	assert.Equal(t, "on", ToString(true, OnOff))
	assert.Equal(t, "off", ToString(false, OnOff))

	type myBool bool
	assert.Equal(t, "true", ToString[myBool](true))
	assert.Equal(t, "false", ToString[myBool](false))
}

func TestToNumeric(t *testing.T) {
	assert.Equal(t, 1, ToNumeric[int](true))
	assert.Equal(t, 0, ToNumeric[int](false))
	assert.Equal(t, uint(1), ToNumeric[uint](true))
	assert.Equal(t, uint(0), ToNumeric[uint](false))

	assert.Equal(t, uint8(1), ToNumeric[uint8](true))
	assert.Equal(t, uint8(0), ToNumeric[uint8](false))

	assert.Equal(t, uint16(1), ToNumeric[uint16](true))
	assert.Equal(t, uint16(0), ToNumeric[uint16](false))

	assert.Equal(t, uint32(1), ToNumeric[uint32](true))
	assert.Equal(t, uint32(0), ToNumeric[uint32](false))

	assert.Equal(t, uint64(1), ToNumeric[uint64](true))
	assert.Equal(t, uint64(0), ToNumeric[uint64](false))

	assert.Equal(t, int8(1), ToNumeric[int8](true))
	assert.Equal(t, int8(0), ToNumeric[int8](false))

	assert.Equal(t, int16(1), ToNumeric[int16](true))
	assert.Equal(t, int16(0), ToNumeric[int16](false))

	assert.Equal(t, int32(1), ToNumeric[int32](true))
	assert.Equal(t, int32(0), ToNumeric[int32](false))

	assert.Equal(t, int64(1), ToNumeric[int64](true))
	assert.Equal(t, int64(0), ToNumeric[int64](false))

	assert.Equal(t, float32(1), ToNumeric[float32](true))
	assert.Equal(t, float32(0), ToNumeric[float32](false))
	assert.Equal(t, float64(1), ToNumeric[float64](true))
	assert.Equal(t, float64(0), ToNumeric[float64](false))
}

func TestToInt(t *testing.T) {
	assert.Equal(t, 1, ToInt[int](true))
	assert.Equal(t, 0, ToInt[int](false))
}

func TestToFloat(t *testing.T) {
	assert.Equal(t, float32(1), ToFloat[float32](true))
	assert.Equal(t, float32(0), ToFloat[float32](false))
	assert.Equal(t, float64(1), ToFloat[float64](true))
	assert.Equal(t, float64(0), ToFloat[float64](false))
}
