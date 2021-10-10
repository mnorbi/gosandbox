package basename

func basename(s string) string {
	lo := 0
	for i := range s {
		if s[i] == '/' {
			lo = i + 1
		}
	}
	hi := len(s)
	for i := hi - 1; i >= lo; i-- {
		if s[i] == '.' {
			hi = i
			break
		}
	}
	return s[lo:hi]
}
