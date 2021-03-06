package inttime

import (
	"fmt"
	"testing"
	"time"
)

// https://www.epochconverter.com/

func TakesSec(s Second) {}
func TakesNano(n Nano)  {}

func TestFoot(t *testing.T) {
	const a = 10
	var n Nano = 1
	var s Second = 2
	fmt.Println(n*a, s*a)

	format := "2006-01-02 15:04:05.000000000"
	time0, err := time.Parse(format, "2021-01-01 00:00:00.000000123")
	if err != nil {
		t.Errorf("setup failure")
	}
	var n0 Nano = FromTime(time0)
	var s0 Second = RoundToSecond(n0)

	if s0 != 1609459200 {
		t.Errorf("bad timestamp %d", s0)
	}

	if n0-ToNano(s0) != 123 {
		t.Errorf("bad difference")
	}
}
