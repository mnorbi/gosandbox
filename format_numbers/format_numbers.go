package format_numbers

import (
	"bytes"
	"strings"
)

// TODO interesting, my thinking was that signature will be:
// func format(int) string
func format(s string) string {
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
