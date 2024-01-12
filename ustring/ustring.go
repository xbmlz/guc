package ustring

import "strings"

// IsEmpty returns true if the string is empty.
func IsEmpty(s string) bool {
	return s == "" || len(s) == 0
}

// IsBlank returns true if the string is empty or contains only white space codepoints.
func IsBlank(s string) bool {
	return IsEmpty(s) || len(strings.TrimSpace(s)) == 0
}

// IsNotEmpty returns true if the string is not empty.
func IsNotEmpty(s string) bool {
	return !IsEmpty(s)
}

// IsNotBlank returns true if the string is not empty and contains at least one non-white space codepoint.
func IsNotBlank(s string) bool {
	return !IsBlank(s)
}
