package signals

import "os"

// Filter filters signal slice in place.
// https://github.com/golang/go/wiki/SliceTricks#filter-in-place
func Filter(a []os.Signal, keep func(x os.Signal) bool) []os.Signal {
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

// Except returns all known signals with exception to specified in input arguments.
func Except(sigs ...os.Signal) []os.Signal {
	keep := func(x os.Signal) bool {
		for i := range sigs {
			if sigs[i] == x {
				return false
			}
		}

		return true
	}

	return Filter(All, keep)
}
