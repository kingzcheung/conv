package str

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestToInt(t *testing.T) {
	assert.Equal(t, 123, ToInt("123", 123))
	assert.Equal(t, 123, ToInt("123b", 123))
	assert.Equal(t, 0, ToInt[string, int]("123b"))
	assert.Equal(t, uint(12234243), ToInt[string, uint]("12234243", uint(123)))
	assert.Equal(t, uint8(12), ToInt[string, uint8]("12", uint8(12)))
	assert.Equal(t, uint16(12), ToInt[string, uint16]("12", uint16(12)))
	assert.Equal(t, uint32(780575), ToInt[string, uint32]("780575", uint32(780575)))
	assert.Equal(t, uint64(370), ToInt[string, uint64]("370", uint64(370)))
	assert.Equal(t, uint64(370), ToInt[string, uint64]("i4P", uint64(370)))
	assert.Equal(t, uint64(0), ToInt[string, uint64]("i4P"))
	assert.Equal(t, int64(-833), ToInt[string, int64]("-833"), int64(-833))
	assert.Equal(t, int32(-833), ToInt[string, int32]("-833bs", int32(-833)))
	assert.Equal(t, int32(-833), ToInt[string, int32]("-833"), int32(-833))
	assert.Equal(t, int32(0), ToInt[string, int32]("-833bs"))
}

func TestToFloat(t *testing.T) {
	assert.Equal(t, 12.34, ToFloat("12.34", 12.34))
	assert.Equal(t, float32(3456464.01), ToFloat[string, float32]("3456464.014"))
	assert.Equal(t, float32(3243.01), ToFloat[string, float32]("3245.014nn", float32(3243.01)))
	assert.Equal(t, float32(0.0), ToFloat[string, float32]("2a4Y4Teh"))
	assert.Equal(t, 0.0, ToFloat[string, float64]("2a4Y4Teh"))
	assert.Equal(t, 12342.789, ToFloat[string, float64]("12342.7890"))
}

func TestToBool(t *testing.T) {
	assert.Equal(t, false, ToBool("abc"))
	assert.Equal(t, true, ToBool("abc", true))
	assert.Equal(t, true, ToBool("true"))
	assert.Equal(t, false, ToBool("false"))

}
