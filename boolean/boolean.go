package boolean

import "github.com/kingzcheung/conv"

type Format int

const (
	TrueFalse Format = iota
	OnOff
	NumberString
)

func ToString[T ~bool](b T, format ...Format) string {
	var trueFt = [3]string{"true", "on", "1"}
	var falseFt = [3]string{"false", "off", "0"}
	defaultFormat := TrueFalse
	if len(format) > 0 {
		defaultFormat = format[0]
	}

	if bool(b) {
		return trueFt[defaultFormat]
	}
	return falseFt[defaultFormat]
}

func ToNumeric[N conv.Numeric](b bool) N {
	if b {
		return N(1)
	}
	return N(0)
}

func ToInt[I conv.Integer](b bool) I {
	return ToNumeric[I](b)
}

func ToFloat[F conv.Float](b bool) F {
	return ToNumeric[F](b)
}
