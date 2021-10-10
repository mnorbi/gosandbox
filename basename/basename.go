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
		// TODO this will break on unicode characters.
		// I think you need to go forward in order to know how to parse unicode
		if s[i] == '.' {
			hi = i
			break
		}
	}
	return s[lo:hi]
}
