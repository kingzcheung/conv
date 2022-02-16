package str

import (
	"github.com/kingzcheung/conv"
	"strconv"
)

func ToInt[T ~string, I conv.Integer](str T, defaultValue ...I) I {
	parseInt, err := strconv.ParseInt(string(str), 10, 64)

	if err != nil {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		} else {
			return 0
		}
	}
	if err != nil {
		return 0
	}
	return I(parseInt)
}

type Float interface {
	~float32 | ~float64
}

func ToFloat[T ~string, F Float](str T, defaultValue ...F) F {
	v, err := strconv.ParseFloat(string(str), 64)
	if err != nil {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		} else {
			return 0.0
		}
	}
	return F(v)
}

func ToBool[T ~string](str T, defaultValue ...bool) bool {
	if str == "true" {
		return true
	}
	if str == "false" {
		return false
	}
	if len(defaultValue) > 0 {
		return defaultValue[0]
	} else {
		return false
	}
}
