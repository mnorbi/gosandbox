package inttime

import "time"

type Nano int64

type Second int64

func FromTime(t time.Time) Nano {
	return Nano(t.UnixNano())
}

func RoundToSecond(n Nano) Second {
	return Second(n / 1e9)
}

func ToNano(s Second) Nano {
	return Nano(s * 1e9)
}
