package utils

// https://github.com/golang/go/wiki/SliceTricks#filter-in-place

// FilterStringSlice filters strings slice in place.
func FilterStringSlice(a []string, keep func(x string) bool) []string {
	n := 0

	for _, x := range a {
		if keep(x) {
			a[n] = x
			n++
		}
	}

	a = a[:n]

	return a
}
