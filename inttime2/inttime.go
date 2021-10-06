package inttime2

import "time"

type Nano int64

func (n Nano) RoundToSecond() Second {
	return Second(int64(n) / 1e9)
}

type Second int64

func (s Second) ToNano() Nano {
	return Nano(s * 1e9)
}

func FromTime(t time.Time) Nano {
	return Nano(t.UnixNano())
}
