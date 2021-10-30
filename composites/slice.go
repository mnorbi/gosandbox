package composites

// RemoveEmpty removes empty strings from sslice.
// Modifies sslice and returns it.
func RemoveEmpty(sslice []string) []string {
	lo := 0
	for hi := 0; hi < len(sslice); hi++ {
		if len(sslice[hi]) > 0 && lo < hi {
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
	len := len(v)
	if k < 0 || k >= len {
		return v
	}
	copy(v[k:], v[k+1:])
	return v[:len-1]
}
