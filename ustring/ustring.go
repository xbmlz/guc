package ustring

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

// IsEmpty returns true if the string is empty.
func IsEmpty(s string) bool {
	return s == "" || len(s) == 0
}

// IsBlank returns true if the string is empty or contains only white space codepoints.
func IsBlank(s string) bool {
	return IsEmpty(s) || len(strings.TrimSpace(s)) == 0
}

// HasEmpty returns true if any of the strings is empty.
func HasEmpty(s ...string) bool {
	if len(s) == 0 {
		return true
	}
	for _, v := range s {
		if IsEmpty(v) {
			return true
		}
	}
	return false
}

// HasBlank returns true if any of the strings is empty or contains only white space codepoints.
func HasBlank(s ...string) bool {
	if len(s) == 0 {
		return true
	}
	for _, v := range s {
		if IsBlank(v) {
			return true
		}
	}
	return false
}

// RemovePrefix removes the prefix from the string.
// e.g. RemovePrefix("abc", "a") => "bc"
func RemovePrefix(s, p string) string {
	if IsEmpty(s) || IsEmpty(p) {
		return s
	}
	if strings.HasPrefix(s, p) {
		return s[len(p):]
	}
	return s
}

// RemoveSuffix removes the suffix from the string.
// e.g. RemoveSuffix("abc", "c") => "ab"
func RemoveSuffix(s, p string) string {
	if IsEmpty(s) || IsEmpty(p) {
		return s
	}
	if strings.HasSuffix(s, p) {
		return s[:len(s)-len(p)]
	}
	return s
}

// Format replaces "{}" with "%v" in the string.
// e.g. Format("abc{}def", "123") => "abc123def"
func Format(s string, args ...interface{}) string {
	res := s
	for _, arg := range args {
		res = strings.Replace(res, "{}", fmt.Sprint(arg), 1)
	}
	return res
}

// Capitalize returns a copy of the string with its first character capitalized and the rest lowercased.
func Capitalize(s string) string {
	if IsEmpty(s) {
		return s
	}
	return strings.ToUpper(s[:1]) + s[1:]
}

// Chomp Removes one newline from end of a String if it's there, otherwise leave it alone. A newline is "\n", "\r", or "\r\n".
func Chomp(s string) string {
	for _, suffix := range []string{"\r\n", "\n", "\r"} {
		if strings.HasSuffix(s, suffix) {
			return strings.TrimSuffix(s, suffix)
		}
	}
	return s
}

// Chop Remove the last character from a String.
func Chop(s string) string {
	if IsEmpty(s) {
		return s
	}
	_, size := utf8.DecodeLastRuneInString(s)
	if strings.HasSuffix(s, "\r\n") {
		return s[:len(s)-2]
	} else if strings.HasSuffix(s, "\n") || strings.HasSuffix(s, "\r") {
		return s[:len(s)-1]
	}
	return s[:len(s)-size]
}
