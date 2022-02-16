package integer

import (
	"github.com/kingzcheung/conv"
	"strconv"
)

func ToString[T conv.Integer](integer T) string {
	return strconv.FormatInt(int64(integer), 10)
}

func ToBool[T conv.Integer](integer T) bool {
	if integer > 0 {
		return true
	}
	return false
}

func ToFloat[T conv.Integer, F ~float32 | ~float64](integer T) F {
	return F(integer)
}
