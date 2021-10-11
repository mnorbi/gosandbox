package format_numbers

import (
	"bytes"
	"strings"
)

func formatNumber(v int) string {
	if v == 0 {
		return "0"
	}
	var buf bytes.Buffer
	for i, j := v, 0; i != 0; {
		buf.WriteRune(rune('0' + abs(i%10)))
		i = i / 10
		j++
		if i != 0 && j%3 == 0 {
			buf.WriteRune(',')
		}
	}
	if v < 0 {
		buf.WriteRune('-')
	}
	return string(reverse(buf.Bytes()))
}
func abs(v int) int {
	if v < 0 {
		return -v
	}
	return v
}
func reverse(b []byte) []byte {
	for lo, hi := 0, len(b)-1; lo < hi; {
		b[lo], b[hi] = b[hi], b[lo]
		lo++
		hi--
	}
	return b
}

// TODO interesting, my thinking was that signature will be:
func formatNumberStr(s string) string {
	s = strings.TrimSpace(s)
	t := s
	var buf bytes.Buffer
	if len(s) > 0 && s[0] == '-' {
		buf.WriteRune('-')
		s = s[1:]
	}
	if len(s) < 4 {
		return t
	}
	i := len(s) % 3
	buf.WriteString(s[0:i])
	for ; i < len(s); i += 3 {
		buf.WriteRune(',')
		buf.WriteString(s[i : i+3])
	}
	return buf.String()
}
