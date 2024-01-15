package uslice

// IsExist checks if a value exists in a slice.
// e.g. IsExist[string]([]string{"a", "b"}, "a") => true
func IsExist[T comparable](s []T, v T) bool {
	for _, e := range s {
		if e == v {
			return true
		}
	}
	return false
}
