package basename

import "unicode/utf8"

func basename(s string) string {
	//TODO: use range
	//TODO: investigate strings.LastIndex
	lo, hi := 0, len(s)
	for i := 0; i < hi; {
		r, size := utf8.DecodeRuneInString(s[i:])
		if r == '/' {
			lo = i + 1
		}
		i += size
	}
	for i := hi; i > lo; {
		r, size := utf8.DecodeLastRuneInString(s[lo:i])
		if r == '.' {
			hi = i - 1
			break
		}
		i -= size
	}
	return s[lo:hi]
}
