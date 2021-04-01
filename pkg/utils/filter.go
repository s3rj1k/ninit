package utils

import (
	"io/fs"
)

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

// IsStringInSlice checks for occurrence of provided string in string slice.
func IsStringInSlice(str string, a []string) bool {
	for i := range a {
		if a[i] == str {
			return true
		}
	}

	return false
}

// FilterDirEntries filters `[]fs.DirEntry` in place by file name.
// (fs.DirEntry contains entries with single depth, meaning no folder recursion).
func FilterDirEntries(a []fs.DirEntry, keep func(x string) bool) []fs.DirEntry {
	n := 0

	for _, x := range a {
		if keep(x.Name()) {
			a[n] = x
			n++
		}
	}

	a = a[:n]

	return a
}
