package composites

// RemoveEmpty removes empty strings from sslice.
// Modifies sslice and returns it.
func RemoveEmpty(sslice []string) []string {
	lo := 0
	for hi := range sslice {
		if sslice[hi] != "" && lo < hi {
			sslice[lo] = sslice[hi]
			sslice[hi] = ""
		}
		if len(sslice[lo]) > 0 {
			lo++
		}

	}
	return sslice[:lo]
}

// Remove the element with index k. Modifies v.
func Remove(v []int, k int) []int {
	if k < 0 || len(v) <= k {
		return v
	}
	copy(v[k:], v[k+1:])
	return v[:len(v)-1]
}
